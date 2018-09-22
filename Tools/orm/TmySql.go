package orm

import (
	"database/sql"
	"strings"
	"fmt"
	//"reflect"
	//"strconv"
)

/*
  MysqlDB对象
*/
type MysqlDB struct{
	*sql.DB
	Params
}
type ParamField struct {
	name string
	value interface{}
}

//type  condition string
//
//const (
//	LESS condition = "<" // value --> 0
//	LESSEQUAL condition = "lte"       // value --> 1
//	LARGE condition = "gt"				// value --> 2
//	LARGEEQUAL condition = "gte"          // value --> 3
//	EQUAL condition = "="
//)


/*
  参数
*/
type Params struct{
	from string  //表名
	fields []string  //select的字段名
	where []ParamField // where条件，暂时只支持and连接
	values []interface{} //查询的参数值
}

/*
  获取数据库对象
*/
func NewDb(driverName, dataSourceName string)(*MysqlDB, error){
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil{
		return nil, err
	}
	mydb := &MysqlDB{DB:db}
	return mydb, nil
}
/*
  插入函数
*/
func (this *MysqlDB)Insert(model interface{}) error{
	strsql, params, tbInfo, err := generateInsertSql(model)
	if err != nil{
		return err
	}
	var result sql.Result
	result, err = this.Exec(strsql, params...)
	if err != nil{
		return err
	}
	setAuto(result, tbInfo)
	return nil
}
/*修改函数*/
func (this *MysqlDB)Update(model interface{}) error{
	strsql, params,  err := generateUpateSql(model)
	if err != nil{
		return err
	}
	_, err = this.Exec(strsql, params...)
	if err != nil{
		return err
	}
	return  nil
}

/*
  删除函数
*/
func (this *MysqlDB)Delete(model interface{}) error{
	strSql, params, err := generateDeleteSql(model)
	if err != nil{
		return err
	}
	_, err = this.Exec(strSql, params...)
	if err != nil{
		return err
	}
	return nil
}

/*设置查询表名*/
func(msqlDB *MysqlDB) From(name string)*MysqlDB{
	msqlDB.from=name
	return msqlDB
}
/*select*/
func (msqlDB *MysqlDB)Select(args ...string)*MysqlDB  {
	for _,v:=range args{
		msqlDB.fields=append(msqlDB.fields,v)
	}
	return  msqlDB
}
/*delete*/


//获取参数
func (this MysqlDB)getValues()[]interface{}{
	return this.values
}
/*where*/
func (msqlDB *MysqlDB) Where(name string,value interface{})*MysqlDB {
	msqlDB.where=append(msqlDB.where,ParamField{name:name,value:value})
	return msqlDB
}

func(msqlDB *MysqlDB) getSelectSql()string{
	var strsql string=" select "
	/*拼接查询参数*/
	if msqlDB.fields !=nil || len(msqlDB.fields)==0 {
		strsql+="*"
	}else {
		for _,v:=range msqlDB.fields{
			strsql+=v+","
		}
		strsql=strings.TrimRight(strsql,",")
	}
	strsql+=" from "+msqlDB.from
	/*拼接查询条件*/
	if msqlDB.where !=nil || len(msqlDB.where)>0 {
		strsql+=" where "
		for _,v:=range msqlDB.where{
			msqlDB.values=append(msqlDB.values,v.value)
			if strings.Contains(v.name,"__"){
				/*条件*/
				nameArgs:=strings.Split(v.name,"__")
				if len(nameArgs)==2 {
					switch nameArgs[1] {
					case "lt":
						strsql += nameArgs[0] + "<? and "
						break
					case "lte":
						strsql += nameArgs[0] + "<=? and "
						break
					case "gt":
						strsql += nameArgs[0] + " > ? and "
						break
					case "gte":
						strsql += nameArgs[0] + ">=? and "
						break
					case "eq":
						strsql += nameArgs[0] + "=? and "
						break
					}
				}else{
					strsql += v.name + "=? and "
				}
			}else {
				strsql += v.name + "=? and "
			}
			strsql = strings.TrimRight(strsql, " and ")
		}
	}
	return strsql
}

/*
  原生Sql查询语句查询
*/
func (this *MysqlDB)Query(sql string, params []interface{}) (*MyRows, error){
	rows, err := this.DB.Query(sql, params...)
	if err != nil{
		return nil, err
	}
	myrows := &MyRows{Rows:rows, Values:make(map[string]interface{})}
	return myrows, nil
}




//执行查询语句，并映射结果到实体切片
func (msqlDB *MysqlDB)Get()([]interface{}, error){
	sql :=  msqlDB.getSelectSql()
	fmt.Println("sql: ", sql)
	vals := msqlDB.getValues()
	fmt.Println("params: ", vals)
	msqlDB.Query(sql,vals)
	rows,err:= msqlDB.Query(sql, vals)
	if err != nil{
		return nil, err
	}
	return rows.To(msqlDB.from)
}

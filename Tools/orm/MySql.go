package orm

import (
	"database/sql"
	//"fmt"
	//"strings"
)

/*
  MysqlDB对象
*/
type MysqlDB struct{
	*sql.DB
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
	strSql, params, tbInfo, err := generateInsertSql(model)
	if err != nil{
		return err
	}
	var result sql.Result
	result, err = this.Exec(strSql, params...)
	if err != nil{
		return err
	}
	setAuto(result, tbInfo)
	return nil
}
/*
  设置自增长字段的值
*/
func setAuto(result sql.Result, tbInfo *Table)(err error){
	defer func(){
		if e := recover(); e != nil{
			err = e.(error)
		}
	}()
	id, err := result.LastInsertId()
	if id == 0{
		return
	}
	if err != nil{
		return
	}
	for _, v := range tbInfo.Fields{
		if v.IsAutoGenerate && v.Value.CanSet(){
			v.Value.SetInt(id)
			break
		}
	}
	return
}
package orm
import (
	"database/sql"
	"reflect"
	"strconv"
)

type MyRows struct{
	* sql.Rows
	Values map[string]interface{} //表字段和值的映射
	ColumnNames []string //表字段名集合
}
//处理返回的结果
func (this *MyRows)Next()bool{
	result:=this.Rows.Next()
	if result{
		if this.ColumnNames==nil||len(this.ColumnNames)==0{
			this.ColumnNames,_=this.Rows.Columns()
		}
		//初始化返回值map
		if this.Values==nil{
			this.Values=make(map[string]interface{})
		}
		//初始化扫描对应数组
		scanArgs:=make([]interface{},len(this.ColumnNames))
		//注意此处是二维数组,因为返回的每个值本身是一个byte数组
		values:=make([][]byte,len(this.ColumnNames))

		for i :=range  values{
			scanArgs[i]=&values[i]
		}
		//将扫描后的值放进返回的map中
		this.Rows.Scan(scanArgs...)
		for i:=0;i<len(this.ColumnNames);i++{
			this.Values[this.ColumnNames[i]]=values[i]
		}
	}
	return  result
}

/*
将数据映射到实体切片
tbname：U对应的数据表名
*/
func (this *MyRows)To(tbname string) ([]interface{},error){
	mi := modelMapping[tbname]
	ti, _ := getTableInfo(mi.Model)
	var models []interface{}
	for this.Next(){
		v := reflect.New(reflect.TypeOf(mi.Model).Elem()).Elem()
		for k, val := range this.Values{
			f := v.FieldByName(ti.TableModelMap[k])
			var strVal string
			if bt, ok := val.([]byte); ok{
				strVal = string(bt)
				switch f.Type().Name(){
				case "int":
					i, _ := strconv.ParseInt(strVal, 10, 64)
					f.SetInt(i)
					break
				case "string":
					f.SetString(strVal)
					break
				}
			}
		}
		models = append(models, v.Interface())
	}
	return models, nil
}
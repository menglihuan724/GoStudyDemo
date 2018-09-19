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

func (this *MyRows)Next()bool{
	return  true
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
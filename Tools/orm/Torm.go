package orm

import    ( "fmt"
	//"database/sql"
	"errors"
	"strings"
	"reflect"
)

/*table entity*/

type Field struct {
	Name string
	IsPrimarykey bool
	IsAutoGenerate bool
	Value reflect.Value
}

type Table struct {
	//表名
	Name string
	//字段映射
	Fields []Field
	//table model 映射
	TableModelMap map[string]string

}

type ModelInfo struct{
	Table
	TbName string
	Model interface{}
}


type TableName string

var typeTableName TableName

var tableNameType reflect.Type=reflect.TypeOf(typeTableName)

var modelMapping map[string]ModelInfo

func RegModel(model interface{})  {
	if  modelMapping ==nil {
		modelMapping=make(map[string]ModelInfo)
	}
	tbInfo,_:=getTableInfo(model)
	modelMapping[tbInfo.Name]=ModelInfo{TbName:tbInfo.Name,Model:model}
}

func getTableInfo(model interface{})(tabinfo *Table,err error) {

	defer func() {
		if e:=recover();err !=nil{
			tabinfo=nil
			err=e.(error)
		}
	}()

	err=nil
	tabinfo=&Table{}
	tabinfo.TableModelMap=make(map[string]string)
	rt:=reflect.TypeOf(model)
	rv:=reflect.ValueOf(model)
	tabinfo.Name=rt.Name()
	if rt.Kind() ==reflect.Ptr{
		rt = rt.Elem()
		rv = rv.Elem()
	}
	for i,j:=0,rt.NumField();i<j;i++{
		rtf:=rt.Field(i)
		rvf:=rv.Field(i)
		if rtf.Type == tableNameType{
			tabinfo.Name = string(rtf.Tag)
			continue
		}
		if rtf.Tag == "-"{
			continue
		}

		var f Field
		if rtf.Tag == ""{
			f = Field{Name:rtf.Name, IsAutoGenerate:false, IsPrimarykey:false, Value:rvf}
			tabinfo.TableModelMap[rtf.Name] = rtf.Name
		}else{
			strTag :=string(rtf.Tag)
			if strings.Index(strTag,":")==-1{
				f = Field{Name:rtf.Name, IsAutoGenerate:false,IsPrimarykey:false, Value:rvf}
				tabinfo.TableModelMap[rtf.Name] = rtf.Name
			}else {
				strName := rtf.Tag.Get("name")
				if strName == ""{
					strName = rtf.Name
				}
				//解析tag中的PK
				isPk := false
				strIspk := rtf.Tag.Get("PK")
				if strIspk == "true"{
					isPk = true
				}
				//解析tag中的auto
				isAuto := false
				strIsauto := rtf.Tag.Get("auto")
				if strIsauto == "true"{
					isAuto = true
				}
				f = Field{Name:strName, IsPrimarykey:isPk, IsAutoGenerate:isAuto, Value:rvf}
				tabinfo.TableModelMap[strName] = rtf.Name
			}
		}
		tabinfo.Fields = append(tabinfo.Fields, f)
	}
	return
}

func generateInsertSql(model interface{})(string, []interface{}, *Table, error){
	//获取表信息
	tbInfo, err := getTableInfo(model)
	if err != nil{
		return "", nil, nil, err
	}
	if len(tbInfo.Fields) == 0 {
		return "", nil, nil, errors.New(tbInfo.Name + "结构体中没有字段")
	}

	//根据字段信息拼Sql语句，以及参数值
	strSql := "insert into " + tbInfo.Name
	strFileds := ""
	strValues := ""
	var params []interface{}
	for _, v := range tbInfo.Fields{
		if v.IsAutoGenerate {
			continue
		}
		strFileds += v.Name + ","
		strValues += "?,"
		params = append(params, v.Value.Interface())
	}
	if strFileds == ""{
		return "", nil, nil, errors.New(tbInfo.Name + "结构体中没有字段，或只有自增字段")
	}
	strFileds = strings.TrimRight(strFileds, ",")
	strValues = strings.TrimRight(strValues, ",")
	strSql += " (" + strFileds + ") values(" + strValues + ")"
	fmt.Println("sql: ",strSql)
	fmt.Println("params: ",params)
	return strSql, params, tbInfo, nil
}

package main

import ("fmt"
	"GoStudyDemo/Tools/orm"
	_ "github.com/go-sql-driver/mysql"
	//"time"
	//"log"
	//"strings"
)

type UserInfo struct{
	TableName orm.TableName "userinfo"
	UserName string `name:"username"`
	Uid int `name:"uid"PK:"true"auto:"true"`
	DepartName string `name:"departname"`
	Created string `name:"created"`
}
func main() {
	ui := UserInfo{UserName:"CHAIN", DepartName:"TEST", Created:"2018-09-16"}
	orm.RegModel(new(UserInfo))
	db, err := orm.NewDb("mysql", "menglihuan:297234@tcp(192.168.42.116:3306)/pwdstore?charset=utf8")
	if err != nil {
		fmt.Println("打开SQL时出错:", err.Error())
		return
	}
	defer db.Close()

	////插入测试
	//err = db.Insert(&ui)
	//if err != nil {
	//	fmt.Println("插入时错误:", err.Error())
	//}

	//查询测试
	res, err := db.From("userinfo").
		Select("username", "departname", "uid").
		Where("username", "CHAIN").Get()
	if err != nil{
		fmt.Println("err: ", err.Error())
	}
	fmt.Println(res)

	ui.UserName = "BBBB"
	err = db.Update(ui)
	if err != nil {
		fmt.Println("修改时错误:", err.Error())
	}
}

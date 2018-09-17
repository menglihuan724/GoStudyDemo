package test

import ("fmt"
	"GoStudyDemo/Tools/orm"
	_ "github.com/go-sql-driver/mysql"
	//"time"
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
	db, err := orm.NewDb("mysql", "menglihuan:297234@tcp(192.168.0.102:3306)/pwdstore?charset=utf8")
	if err != nil {
		fmt.Println("打开SQL时出错:", err.Error())
		return
	}
	defer db.Close()

	//插入测试
	err = db.Insert(&ui)
	if err != nil {
		fmt.Println("插入时错误:", err.Error())
	}
}

package Tools

import (

	"github.com/Luxurioust/excelize"

	"database/sql"
	"fmt"
	"os"
	"strconv"
	"../../myGo/tool/entity"
)

func main() {
	 admin :=new(entity.Admin)
	//连接数据库
	db, err := sql.Open("mysql", "root:297234@tcp(localhost:3306)/weixin?charset=utf8")
	checkErr(err)
	admins, err := db.Query("SELECT admin_id,username FROM admin")
	checkErr(err)
	for admins.Next() {
		var admin_id int
		var username string
		err = admins.Scan(&admin_id, &username)
		checkErr(err)
		fmt.Println(admin_id)
		fmt.Println(username)
	}
	xlsx, err := excelize.OpenFile("D:/myworkspace/ZdsyzbImport/tool/test.xlsx")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Get value from cell by given sheet index and axis.
	//cell := xlsx.GetCellValue("Sheet1", "B2")
	//fmt.Println(cell)
	// Get sheet index.
	index := xlsx.GetSheetIndex("Sheet1")
	// Get all the rows in a sheet.
	rows := xlsx.GetRows("Sheet" + strconv.Itoa(index))
	for i, row := range rows {
		if(i>0){
			for j, colCell := range row {
				switch j{
				case 0:admin.admin_id,err=strconv.Atoi(colCell)
				case 1:admin.username=colCell
				case 2:admin.realname=colCell
				case 3:admin.pwd=colCell
				}
				fmt.Print("admin=",admin)

			}
			//插入数据
			stmt, err := db.Prepare("INSERT  INTO  admin  (admin_id,username,realname,pwd) VALUES (?,?,?,?)")

			checkErr(err)

			fmt.Print("admin2=",admin)

			res,err :=stmt.Exec(strconv.Itoa(admin.admin_id),admin.username,admin.realname,admin.pwd)

			id, err := res.LastInsertId()
			checkErr(err)
			fmt.Println(id)

		}

	}
	db.Close()
}

//func checkErr(err error) {
//	if err != nil {
//		panic(err)
//	}
//}



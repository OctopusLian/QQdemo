package main

import (
	"fmt"
	"QQdemo/Server/DBServ/DBManger"
)
var(
	dbhostsip = "127.0.0.1:3306"  //IP地址
	dbusername = "root"  //数据库的用户名
	dbpassword = "123456"  //登录数据库的密码
	dbname = "user"  //数据库表名
)



func main() {
	var i int
	for{
		fmt.Println("1-开启数据库，2-插入，3-查询，4-更新，5-删除，6-关闭数据库")

		fmt.Println("请输入序号")
		fmt.Scan(&i)
		//var userid int
		switch i {
		case 1:
			DBManger.DBstart()
		case 2:
			DBManger.DBinsert()
		case 3:
			DBManger.DBselect()
		case 4:
			DBManger.DBupdate(DBManger.Userid)
		case 5:
			DBManger.DBdelete(DBManger.Userid)
		case 6:
			DBManger.DBclose()
		}
	}
}


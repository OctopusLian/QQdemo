package DBManger

import (
	"database/sql"
	"fmt"
    _ "DBAndProtobuf_Gotest/github.com/go-sql-driver/mysql"
)

var db *sql.DB

var Username,Password string
var Userid int

func DBstart() {
	db,_= sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/qqdemo?charset=utf8") //qqdemo是我新建立的数据库名

	fmt.Println("qqdemo database connect succcessed !")
}


func CheckErr(err error){
	if err != nil{
		panic(err)
		//fmt.Println(err)
	}
}
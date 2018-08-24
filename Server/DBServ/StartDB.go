package DBServ

import (
	"database/sql"
	"fmt"
)

var db *sql.DB
var err error

func init(){
	db,err = sql.Open("MySQL","root:123456@tcp(127.0.0.1:3036)/qqdemo?charset=utf8")
	if err != nil{
		checkErr("[DBServ] Connect",err)
	}
}

func checkErr(s string,err error){
	if err != nil{
		fmt.Println(err)
	}
}

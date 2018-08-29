package DBManger

import "fmt"

func DBclose(){
	//关闭数据库
	db.Close()
	fmt.Println("qqdemo database closed")
}

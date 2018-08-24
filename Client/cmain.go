package main

import (
	"fmt"
	"net"
)

var Name string = ""
var Password string = ""

func main(){
	fmt.Println("请输入ip和端口号：")
	var addr_port string
	fmt.Scan(&addr_port)
	conn,err := net.Dial("tcp","addr_port")
	if err != nil{
		fmt.Println("failed")
	}
	defer conn.Close()
	fmt.Println("connect successed !")

	for {
		fmt.Println("1登录；2获取房间列表；3进入游戏")
		var i int
		fmt.Println("测试")
		fmt.Scan(&i)
		switch i {
		case 1:
			fmt.Printf("scanf your userid:")
			var userid  string
			fmt.Scan(&userid)


		}
	}
}

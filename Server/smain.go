package main

import (
	"fmt"
	"QQdemo/Server/ServerNet"
	"net"
	"QQdemo/Server/helper"
)

var Addr_port string
var Server_name string

/*
func GetNodeMap() map[string]Node {
	return nodeMap
}
*/

func main(){
	fmt.Println("请输入本地ip和端口号：")
	fmt.Scan(&Addr_port)
	fmt.Println("请输入游戏服务器名字：")
	fmt.Scan(&Server_name)

	go func() {
		for{
			var message string
			fmt.Scan(&message)
			switch message {
			case "print":

			case "close":
				ServerNet.ServerClose()

			}
		}
	}()

	ServerNet.ServerClose()

}

func Send(conn net.Conn,msg []byte){
	lenBytes := helper.NumToByte(int32(len(msg)))
	sendBuff := helper.ConcatByte(lenBytes,msg)
	conn.Write(sendBuff)
}

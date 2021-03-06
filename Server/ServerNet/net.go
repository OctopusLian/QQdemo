package ServerNet

import (
	"net"
	"QQdemo/Server/helper"
	"fmt"
)

var listen_socket net.Listener
var err error

//心跳时间
var heartBeatTime int64 = 60

//用户列表
//var nodeMap map[string]Node = make(map[string]Node)
//var lockMap sync.Mutex

//开启服务
func ServerStart(addr_port string){
	listen_socket,err = net.Listen("tcp",addr_port)
	if err != nil{
		fmt.Println("server start fail")
	}
	fmt.Println("server is start on :",addr_port)

	//添加节点
	//cNode := newNode(conn)

	//lockMap.Lock()
	//AddNodeMap(conn.R)
}

//关闭服务
func ServerClose(){
	listen_socket.Close()
}

//发送消息 加上包长
func Send(conn net.Conn, msg []byte) {
	lenBytes := helper.NumToByte(int32(len(msg)))
	sendBuff := helper.ConcatByte(lenBytes, msg)
	fmt.Println("Send:", sendBuff)
	conn.Write(sendBuff)
}






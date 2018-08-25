package ServerNet

import (
	"QQdemo/Server/proto"
	proto2 "QQdemo/github.com/golang/protobuf/proto"
	"fmt"
	"QQdemo/Server/helper"
)

//登录
func OnLogin(){
	LoginProto := &proto.Login_ToS{}
	err := proto.Unmarshal(protobyte,LoginProto)
	if err != nil{
		fmt.Println(err)
		return
	}
	var user_id string = LoginProto.GetUid()
	fmt.Println("OnLogin :",user_id)

	LoginToC := &proto.Login_ToC{}
	var cmd []byte = helper.NumToByte(int32(101))

	if !KickOff(user_id){
		LoginToC.IsOk = proto.Bool(false)
		//发踢下线消息
		buffer,_ := proto2.Marshal(LoginToC)
		SendMsg := helper.ConcatByte(cmd,buffer)
		Send(,SendMsg)
		return
	}

}

//踢下线
func KickOff(userid string) bool{
	for k,v := range userlist{

	}
}

//下线
func Logout(){

}


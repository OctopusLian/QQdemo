package ServerNet

import (
	"QQdemo/Server/proto"
	proto2 "QQdemo/github.com/golang/protobuf/proto"
	"fmt"
)

//登录
func OnLogin(){
	LoginProto := &proto.Login_ToS{}
	err := proto2.Unmarshal()
	if err != nil{
		fmt.Println(err)
		return
	}
	var user_id string = LoginProto.GetUid()
	fmt.Println("OnLogin :",user_id)

}

//更新


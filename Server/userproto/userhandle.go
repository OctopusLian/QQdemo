package userproto

var Handles map[int]string = make(map[int]string)

func init(){
	Handles[1] = "HeartBeat"

	Handles[101] = "LoginToS"
	Handles[102] = "LoginToC"
	Handles[103] = "RegisterToS"
	Handles[104] = "RegisterToC"
	Handles[105] = "GetRoomListToS"
	Handles[106] = "GetRoomListToC"
}

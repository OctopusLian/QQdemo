package DBServ

//用户注册
func Register(user_id string) bool{
	//将用户id插入数据库
	stmt,_ := db.Prepare("Insert user set user_id = ?")
	_,err := stmt.Exec(user_id)
	if err != nil{
		return false
		checkErr("[DBServ] Insert Successed !",err)
	}
	return true
}

package DBServ

//用户注册
func Register(userid string) bool{
	//将用户id插入数据库
	stmt,_ := db.Prepare("Insert user set userid = ?")  //使用Prepare获得stmt
	_,err := stmt.Exec(userid)  //调用Exec添加
	if err != nil{
		return false
		checkErr("Register failed !",err)
	}
	return true
}

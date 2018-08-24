package DBServ

//查询用户id是否已经存在
func CanRegister(user_id string) bool{
	//查询数据
	rows,err := db.Query("Select * from user where user_id = ?",user_id)
	checkErr("[DBServ] Select",err)
	if rows.Next(){
		return true
	}
	return false
}

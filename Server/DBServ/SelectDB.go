package DBServ

import "fmt"

//查询用户id是否已经存在
func CanRegister(userid string) bool{
	//查询数据
	rows,err := db.Query("Select * from user where userid = ?",userid)
	checkErr("Select failed !",err)
	if rows.Next(){
		for rows.Next(){
			var userid string
			var username string
			var password string
			se := rows.Scan(&userid, &username, &password)
			checkErr("Select end and failed !",se)
			fmt.Println(userid)
			fmt.Println(username)
			fmt.Println(password)
		}
		return true
	}
	return false
}

package DBManger

import "fmt"

func DBupdate(userid int){
	//更新数据
	fmt.Println("请输入要查询的用户ID：")
	fmt.Scan(&Userid)
	fmt.Println("请输入要更新的内容：a-名字；b-密码：")
	var s string
	fmt.Scan(&s)
	switch s {
	case "a":
		fmt.Println("请输入要更改用户ID为%d的名字",Userid)
		fmt.Scan(&Username)
		stmt,err := db.Prepare("UPDATE user SET username=? where userid=?")
		CheckErr(err)

		fmt.Println("用户%d更新后的名字为：",Userid)
		fmt.Println(Username)
		res,err := stmt.Exec(Username,Userid)
		affect,err := res.RowsAffected()
		fmt.Println(affect)
	case "b":
		fmt.Println("请输入要更改用户ID为%d的密码",Userid)
		fmt.Scan(&Password)
		stmt,err := db.Prepare("UPDATE user SET password=? where userid=?")
		CheckErr(err)

		fmt.Println("用户%d更新后的名字为：",Userid)
		fmt.Println(Password)
		res,err := stmt.Exec(Password,Userid)
		affect,err := res.RowsAffected()
		fmt.Println(affect)
	default:

	}









}

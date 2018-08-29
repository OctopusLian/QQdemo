package DBManger

import(
	"fmt"
)

func DBinsert(){

	fmt.Println("请输入注册用户ID：")
	fmt.Scan(&Userid)
	fmt.Println("请输入注册的用户昵称：")
	fmt.Scan(&Username)
	fmt.Println("请输入注册账号的密码：")
	fmt.Scan(&Password)

	//插入数据
	stmt,err := db.Prepare("INSERT user SET userid=?,username=?,password=?")
	CheckErr(err)

	res,err := stmt.Exec(Userid,Username,Password)
	CheckErr(err)

	id,err := res.LastInsertId()

	fmt.Println(id)
}


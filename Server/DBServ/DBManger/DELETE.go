package DBManger

import "fmt"

func DBdelete(userid int){
	//删除数据

	fmt.Println("请输入要删除的用户ID：")
	fmt.Scan(&Userid)
	stmt,err := db.Prepare("DELETE from userif where userid=?")
	CheckErr(err)

	res,err := stmt.Exec(Userid)
	CheckErr(err)

	affect,err := res.RowsAffected()
	CheckErr(err)

	fmt.Println(affect)
}

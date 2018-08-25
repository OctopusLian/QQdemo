package DBServ

import "fmt"

//删除操作
func Dele(userid string){
	stmt,err := db.Prepare("delete from user where userid=?")
	checkErr("Delete failed！",err)


	res,err := stmt.Exec(userid)
	checkErr("Delete failed !",err)

	affect,err := res.RowsAffected()
	checkErr("Delete data failed !",err)

	fmt.Println(affect)

	db.Close()
}
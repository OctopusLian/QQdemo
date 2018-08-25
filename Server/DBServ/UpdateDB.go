package DBServ

import "fmt"

func Upda(userid string){
	stmt,err := db.Prepare("update user set username=? where uid=?")
	checkErr("Update failed !",err)



	res,err := stmt.Exec(userid)
	checkErr("Update data failed !",err)

	affect,err := res.RowsAffected()
	checkErr("Updata rows failed !",err)

	fmt.Println(affect)
}

package DBManger

import "fmt"

func DBselect(){
	//查询数据
	rows,err := db.Query("SELECT * FROM user")
	CheckErr(err)

	for rows.Next(){

		err = rows.Scan(&Userid,&Username,&Password)
		CheckErr(err)
		fmt.Println(Userid)
		fmt.Println(Username)
		fmt.Println(Password)
	}
}

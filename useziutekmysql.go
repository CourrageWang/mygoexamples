package main

import (
	"fmt"

	"github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/native"
)

func main() {
	getUsrIfno("allen")
	if ok:=test2("1");!ok {
		fmt.Println("bu ok")
	}

}
func getUsrIfno(userAccount string) {
	db2 := mysql.New("tcp", "", "127.0.0.1:3306", "root", "12345678", "srun1")
	err := db2.Connect()
	if err != nil {
		fmt.Println(err)
	}
	sqlString := fmt.Sprintf("SELECT age FROM teacher where name = '%s'", userAccount)

	fmt.Println(sqlString)
	rows, _, err2 := db2.Query(sqlString)

	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(rows[0].Str(0))

}

func test2(sty string) bool {
	if sty != "" {
		if sty == "1" {
			return true
		}
		return false
	}
	return false

}

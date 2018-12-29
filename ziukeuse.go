package main

import (
	"github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/native"
	"git_bak/imc/src/floger"
	"github.com/ziutek/mymysql/autorc"
	"strings"
	"testing"
)

var (
	conn   = []string{"tcp", "", "127.0.0.1:3306"}
	user   = "root"
	passwd = "12345678"
	dbname = "srun1"
)

func getAdmin() {
	db := mysql.New("tcp", "", "127.0.0.1:3306", "root", "12345678", "srun1")
	err := db.Connect()
	if err != nil {
		panic(err)
	}
	sqlstring := "INSERT INTO teacher(`name`,address,age,`year`)VALUES('王五','西安市长安区',17,'2018-10-08')"

	rows, res, err := db.Query(sqlstring)
	floger.Info("err= ", err)

	if err != nil {
		floger.Error("err is", err)
	}
	//if err != nil {
	//	panic(err)
	//}
	if len(rows) < 1 {
		floger.Error("rows error", len(rows))
	}

	floger.Info(res)
	//first := res.Map("username")
	//second := res.Map("password")
	//username, password := row.Str(first), row.Str(second)
	//return username, password
}
func checkErr(t *testing.T, err error, exp_err error) {
	if err != exp_err {
		if exp_err == nil {
			t.Fatalf("Error: %v", err)
		} else {
			t.Fatalf("Error: %v\nExpected error: %v", err, exp_err)
		}
	}
}
func main() {
	//getAdmin()

	c := autorc.New(conn[0], conn[1], conn[2], user, passwd)
	sqlstring := "INSERT INTO teacher(`name`,address,age,`year`)VALUES('周六','西安市长安区',17,'2018-10-08')"
	// Register initialisation commands
	c.Register("set names utf8")
	// my is in unconnected state
	err2 := c.Use(dbname)

	if err2!=nil {
		floger.Error("use db err",err2)
	}
	_, _, err := c.Query(sqlstring)

	if err != nil {
		floger.Error(sqlstring, err)
		if strings.Index(err.Error(), "Duplicate") < 0 {
			floger.Error(err)
		}
	}

}

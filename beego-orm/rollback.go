package main

/**
  mysql 回滚
 */
import (
	_ "github.com/go-sql-driver/mysql" //初始化一个mysql驱动，必须
	"github.com/jmoiron/sqlx"
	"fmt"
)

var Db *sqlx.DB

func init() {
	//"mysql"指定数据库类型，  /test指定打开的数据库  root:123 冒号隔开密码 root账号 123密码
	database, err := sqlx.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/srun1")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = database
}

func main() {
	conn, err := Db.Begin() //开启事务。
	defer Db.Close()
	if err != nil {
		fmt.Println("err.", err)
	}
	r, err2 := conn.Exec("INSERT INTO `teacher` (`name`,`address`,`age`,`year`)VALUES (?,?,?,?)", "walker", "中国", 18, "2018-11-13")
	if err2 != nil {
		fmt.Println("err2", err2)
		return
	}
	id, err3 := r.LastInsertId()

	if err3 != nil {
		fmt.Println("exec failed...")
		conn.Rollback() // 回滚

	}
	fmt.Println("insert ok..", id)
	conn.Commit() //  提交事务

}

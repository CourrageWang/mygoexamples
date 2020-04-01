package mode

import (
	"database/sql"
)

var NewDB *sql.DB

// 连接数据库
func InitnewDB(dataBaseName string) {
	var err error
	NewDB, err = sql.Open("mysql", dataBaseName)
	if err != nil {
		panic(err)
	}
}
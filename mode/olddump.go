package mode

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var oldDB *sql.DB

// 连接数据库
func InitoldDB(dataBaseName string) {
	var err error
	oldDB, err = sql.Open("mysql", dataBaseName)
	if err != nil {
		panic(err)
	}
}

// 组合车场未关联sql  [types为车场类型，0 为new1为old]
func composeCleanCarParkSql(date, dbCloum string) (baseSqlStr string) {
	baseSqlStr = "DELETE FROM t_park_info WHERE NOT EXISTS ( SELECT 1 FROM t_park_channel tpc WHERE tpc.park_code = t_park_info.park_code ) AND " +
		dbCloum + " <= ( SELECT UNIX_TIMESTAMP( DATE_FORMAT( DATE_SUB( NOW( ), INTERVAL " + date +
		" DAY ), '%Y-%m-%d %H:%i:%S' ) ) )"
	return
}

// 组合集团未关联sql
func composeCleanCompanySql(date string) (baseSql string) {

	return
}

// 15天未关联智能设备的车场
func CleanParks15DaysNotBindARM(duration string) {
	/**
	  思路 ： 从park_info表中查找出15天以前注册的车场的park_code，并在park_channel表中查找对应park_code的记录是否存在，
	        如果不存在 说明该车场未绑定只是注册而已 ，则从park_info表中删除该park_code对应的记录。【老库新库相同】
	*/
	// 组合sql
	parkNoRelaOldStr := composeCleanCarParkSql(duration, "t_park_info.create_time")
	fmt.Println("删除老库sql 语句为：", parkNoRelaOldStr)

	rowsOld, err := oldDB.Exec(parkNoRelaOldStr)
	if err != nil {
		panic(err.Error())
	}

	defer oldDB.Close()

	sizeOld, _ := rowsOld.RowsAffected()

	if err == nil {
		fmt.Println("清除老库数据成功,共清除数据条数为：", sizeOld)
	} else {
		fmt.Println("清除老库数据失败", err.Error())
	}

	parkNoRelaNewStr := composeCleanCarParkSql(duration, "t_park_info.createon")
	fmt.Println("删除新库sql 语句为：", parkNoRelaNewStr)

	//step2 删除新库中的脏数据
	rowsNew, err2 := NewDB.Exec(parkNoRelaNewStr)

	defer NewDB.Close()
	sizeNew, _ := rowsNew.RowsAffected()
	if err2 == nil {
		fmt.Println("清除新库数据成功,共清除数据条数为：", sizeNew)
	} else {
		fmt.Println("清除新库失败", err2.Error())
	}
}

//集团未关联车场15天
func CleanGroup15DaysNotBindParks(duration string) {

}

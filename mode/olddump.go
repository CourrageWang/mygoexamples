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

// 组合集团未关联sql types【0为新库，1为老库】
func composeCleanCompanySql(date string, types int) string {

	var baseSql string
	if types == 0 {
		baseSql = " DELETE FROM n_group WHERE NOT EXISTS ( SELECT 1 FROM n_park tpc WHERE tpc.group_code = n_group.id) AND UNIX_TIMESTAMP(n_group.create_at) <= ( SELECT UNIX_TIMESTAMP( DATE_FORMAT( DATE_SUB( NOW( ), INTERVAL " +
			date +
			" DAY ), '%Y-%m-%d %H:%i:%S' ) ) )"
	} else {
		baseSql = "DELETE FROM t_company_info WHERE NOT EXISTS ( SELECT 1 FROM t_park_info tpi WHERE tpi.cid = t_company_info.id) AND t_company_info.createon <= ( SELECT UNIX_TIMESTAMP( DATE_FORMAT( DATE_SUB( NOW( ), INTERVAL " +
			date +
			" DAY ), '%Y-%m-%d %H:%i:%S' ) ) )"
	}

	return baseSql
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

	sizeOld, _ := rowsOld.RowsAffected()

	if err == nil {
		fmt.Println("清除老库车场未关联数据成功,共清除数据条数为：", sizeOld)
	} else {
		fmt.Println("清除老库数据失败", err.Error())
	}

	parkNoRelaNewStr := composeCleanCarParkSql(duration, "t_park_info.createon")
	fmt.Println("删除新库sql 语句为：", parkNoRelaNewStr)

	//step2 删除新库中的脏数据
	rowsNew, err2 := NewDB.Exec(parkNoRelaNewStr)

	//defer NewDB.Close()
	sizeNew, _ := rowsNew.RowsAffected()
	if err2 == nil {
		fmt.Println("清除新库数车场未关联数据成功,共清除数据条数为：", sizeNew)
	} else {
		fmt.Println("清除新库失败", err2.Error())
	}
}

//集团未关联车场15天
func CleanGroup15DaysNotBindParks(duration string) {

	//组合sql
	companyNoRelaOldStr := composeCleanCompanySql(duration, 1)
	fmt.Println("老库清除集团未关联超过", duration, "天的sql语句为:", companyNoRelaOldStr)

	rowsOld, err := oldDB.Exec(companyNoRelaOldStr)

	if err != nil {
		panic(err.Error())
	}

	defer oldDB.Close()

	sizeOld, _ := rowsOld.RowsAffected()

	if err == nil {
		fmt.Println("清除老库集团未关联数据成功,共清除数据条数为：", sizeOld)
	} else {
		fmt.Println("清除老库集团未关联数据失败", err.Error())
	}

	//清除新库数据

	companyNoRelaNewStr := composeCleanCompanySql(duration, 0)

	fmt.Println("新库清除集团未关联超过", duration, "天的sql语句为:", companyNoRelaNewStr)

	rowsNew, err2 := NewDB.Exec(companyNoRelaNewStr)

	defer NewDB.Close()

	sizeNew, _ := rowsNew.RowsAffected()
	if err2 == nil {
		fmt.Println("清除新库集团未关联数据成功,共清除数据条数为：", sizeNew)
	} else {
		fmt.Println("清除新库集团未关联失败", err2.Error())
	}
}

package main

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

/**
    beego orm 增删改查
 */
const (
	DRIVERNAME    = "mysql"
	DATASOURCE    = "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8&loc=Asia%2FShanghai"
	MAX_IDLE_CONN = 5
	MAX_OPEN_CONN = 30
)

type User struct {
	Id    int
	Name  string
	Email string
	Age   int
	//isActive bool
}

func (u *User) TableName() string { // 数据库中表名称
	return "t_user"
}
func registerDB() {

	orm.Debug = true
	orm.RegisterDataBase("default", DRIVERNAME, DATASOURCE, MAX_IDLE_CONN, MAX_OPEN_CONN)
	orm.RegisterModel(new(User))
}

// 插入数据
func createUsers() {
	users := []User{
		User{Name: "admin1", Email: "18829290974@163.com", Age: 12},
		User{Name: "admin2", Email: "18829290975@163.com", Age: 13},
		User{Name: "admin2", Email: "18829290976@163.com", Age: 14},
		User{Name: "admin4", Email: "18829290977@163.com", Age: 15},
	}
	o := orm.NewOrm()
	if successNum, err := o.InsertMulti(len(users), users); err != nil {
		fmt.Println("insert fail ...", err)
	} else {
		fmt.Printf("sucess inserted %d datas", successNum)
	}
}

//  查找
func listUsers() {
	var users []User
	orm.NewOrm().QueryTable("t_user").All(&users)
	for _, user := range users {
		fmt.Println(user)
	}
}

// 统计用户的个数

func ountUser() {
	cont, _ := orm.NewOrm().QueryTable("t_user").Count()
	fmt.Printf("all user is %d ", cont)
}

// 根据条件查询用户
func getUser() {
	var user User
	err := orm.NewOrm().QueryTable("t_user").Filter("Id", 3).One(&user)
	if err == nil {
		fmt.Println(user.Name)
	}

}

// 使用 limit offset字段

func limitoffersetUser() {
	var users [] User
	_, err := orm.NewOrm().QueryTable("t_user").Limit(1, 3).OrderBy("id").All(&users)
	if err == nil {
		for _, u := range users {
			fmt.Println(u.Name)
		}
	}

}

//  删除用户
func delUse() {

	num, err := orm.NewOrm().QueryTable("t_user").Filter("id", 2).Delete()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(num)

}

// 更新用户

func updateUser() {

	num, err := orm.NewOrm().QueryTable("t_user").
		Filter("name", "admin1").Update(
		orm.Params{
			"name": "wangyongqi",
		})
	fmt.Printf("%d   %s", num, err)

}

func main() {
	registerDB()
	//createUsers()
	//listUsers()
	//ountUser()
	//getUser()
	//limitoffersetUser()
	//delUse()
	updateUser()
	beego.Run()

}

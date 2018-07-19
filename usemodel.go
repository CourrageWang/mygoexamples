package main

import (
	"test/model"
)

//db, err := sql.Open("mysql",
//        "user:password@tcp(127.0.0.1:3306)/hello")
func main() {
	model.InitDB("user:password@tcp(127.0.0.1:3306)/hello")

}

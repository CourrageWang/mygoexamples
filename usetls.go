package main

import (
	"fmt"
)

func main() {
	userMaps := make([]map[string]string, 0)
	userMap := make(map[string]string)
	for i := 0; i < 10; i++ {
		userMap["userAccount"] = "04131148"
		userMaps = append(userMaps, userMap)
	}
	userAccount :=""
	for _, V := range userMaps {
		//fmt.Println(V["userAccount"])
		 userAccount += "'"+V["userAccount"]+"'"+","
	}
	userAccount = userAccount[:len(userAccount)-1]
	//fmt.Println(userAccount)
    finalSql :=fmt.Sprintf("select `user_id` from `users` where `user_name`in(%s);",userAccount)
	fmt.Println(finalSql)
}

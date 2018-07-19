package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "192.168.0.195:16380")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
// srun3/
	defer c.Close()
	c.Do("Auth","srun_3000@redis")
	Id, err := redis.String(c.Do("GET","key:rad_online:ip:192.168.0.105:0"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("id ", Id)
}

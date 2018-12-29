package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"srun4kAuthIntf/src/floger"
)

func main() {
	c, err := redis.Dial("tcp", "192.168.0.195:16382")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	c.Do("Auth", "srun_3000@redis")
	re, err2 := redis.String(c.Do("Lrange", "lst", "0", "-1"))

	if err2 != nil {
		floger.Error("err2", err2)
	}
	fmt.Println(re)

}

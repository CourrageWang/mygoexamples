package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	result := make([]interface{},0)
	j := 0
	c, err := redis.Dial("tcp", "192.168.0.195:16382")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	c.Do("Auth", "******")
	//list  hash key  list:users*
	//key:interface:*
	//keys, err := redis.Values(c.Do("keys", "list:users:*"))
	keys ,err := redis.Values(c.Do("keys","key:interface:*"))
	if err != nil {
		fmt.Println(err)
	}
	c.Close()
	for _, e := range keys {
		j++ // 计数器
		result = append(result, string(e.([]byte)))
		if len(result) == 500 { // 当查询出ldap的数据大于500条
			wg.Add(1)
			tmp := result
			go func() {
				defer wg.Done()
				SearchResultHandler(tmp)
			}()
			result = nil
		}
		if j == len(keys) { //（总条数-500*处理次数)剩余的条数。
			wg.Add(1)
			go func() {
				defer wg.Done()
				SearchResultHandler(result)
			}()
		}
	}
	wg.Wait()
	fmt.Println("over")
}

func SearchResultHandler(re []interface{})  {
	c, err := redis.Dial("tcp", "192.168.0.195:16382")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	c.Do("Auth", "srun_3000@redis")
	defer c.Close()
	c.Do("del",re...)
}






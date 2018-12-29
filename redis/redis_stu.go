package main

import (
	"github.com/go-redis/redis"
	"floger"
	"time"
	"sync"
	"fmt"
)

/**
   "gopkg.in/redis.v4"  https://github.com/go-redis/redis
 */

var client *redis.Client
//初始化客户端
func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:8765",
		Password: "*****",
		DB:       0,
		PoolSize: 5,
	})
	pong, err := client.Ping().Result()
	if err != nil {
		floger.Error(err)
	}
	floger.Info("redis client response", pong)
}

// String 操作
func StringOption(cli *redis.Client) {
	defer cli.Close()
	//  redis 设置过期时间，（客户端可以以毫秒或者秒为时间，为数据库中的某个键设置生存时间，在经过指定的时间后， 服务器会自动删除过期的键）
	err := cli.Set("age", "test", 0).Err() //没有设置过期时间
	if err != nil {
		floger.Error("err", err)
	}

	result, err2 := cli.Get("age").Result()
	if err2 != nil {
		floger.Error("err2", err2)
	}
	floger.Info("age:", result)

	// 设置过期时间位5秒
	err3 := cli.Set("year", "1992", 5*time.Second).Err()
	if err3 != nil {
		floger.Error("err3", err3)
	}
	cli.Incr("year") // 自增
	cli.Incr("year") // 自增
	cli.Decr("year") //自减

	re, err4 := cli.Get("year").Result()
	if err4 != nil {
		floger.Error(err4)
	}
	floger.Info("age:", re)
}

//  list 操作
func ListOption(cli *redis.Client) {
	defer cli.Close()
	keys := make([]string, 0)
	keys = append(keys, "apple", "banana", "lem") // 批量插list入数据 在尾部追加
	err := cli.RPush("fruit", keys).Err()
	lErr := cli.LPush("fruit", "head..").Err()
	if err != nil {
		floger.Error("RPush err", err)
	}
	floger.Info("RPush suss")
	if lErr != nil {
		floger.Error("lErr err", lErr)
	}
	floger.Info("LPush suss")
	//查询所有的数据
	res, err1 := cli.LRange("fruit", 0, -1).Result()
	if err1 != nil {
		floger.Error("err1..", err1)
	}
	for index, value := range res {
		floger.Info(index, value)
	}

	/*cli.LPop("fruit") // 返回病删除fruit中的首元素
	cli.RPop("fruit") // 返回病删除fruit中的尾元素*/

}

//set 使用方式
func SetOption(cl *redis.Client) {
	defer cl.Close()
	keys := make([]string, 0)
	keys = append(keys, "spr", "win", "sum", "alt")
	err := cl.SAdd("set:session", keys).Err() // set 批量插入
	err2 := cl.SAdd("set:session2", "win", "sum").Err()
	if err != nil {
		floger.Error("SAdd err", err)
	}
	if err2 != nil {
		floger.Error("err2", err2)
	}

	//  判断元素是否在集合中
	exit, err3 := cl.SIsMember("set:session", "alt").Result()
	if err3 != nil {
		floger.Error("err3", err3)
	}
	floger.Info(exit)

	//求元素的交际

	names, err4 := cl.SInter("set:session", "set:session2").Result()
	if err4 != nil {
		floger.Error("err4", err4)
	}
	floger.Info("交集为:", names)

	//  获取指定集合
	results, err5 := cl.SMembers("set:session").Result()
	if err5 != nil {
		floger.Error("err5", err5)
	}
	floger.Info("all results is :", results)
}

// hash 操作
func HashOption(cli *redis.Client) {
	defer cli.Close()
	/*cli.HSet("user", "name", "SRUN")
	cli.HSet("user", "age", 18)

	// 批量插入
	mInfos := make(map[string]interface{})
	mInfos["name"] = "hmset"
	mInfos["age"] = 12
	cli.HMSet("hmset", mInfos)*/
	// 批量获取
	fileds, err := cli.HMGet("user", "name", "age").Result()
	if err != nil {
		floger.Error("err", err)
	}
	floger.Info("fileds:", fileds)

	// 获取user中字段的个数
	count, err2 := cli.HLen("user").Result()
	if err2 != nil {
		floger.Info(err2)
	}
	floger.Info("fields has ", count)

	// 删除user的age字段
	cli.HDel("user", "age")

	res, err3 := cli.HGet("user", "age").Result()
	if err3 != nil {
		floger.Error("err3", err3)
	}
	floger.Info("res is :", res)
}

// 连接池

func Connectpool(clie *redis.Client) {

	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				client.Set(fmt.Sprintf("name%d", j), fmt.Sprintf("xys%d", j), 0).Err()
				client.Get(fmt.Sprintf("name%d", j)).Result()
			}

			floger.Info("PoolStats, TotalCons", client.PoolStats().TotalConns)
		}()
	}
	wg.Wait()
}

func main() {

	//ListOption(client)
	//SetOption(client)
	//HashOption(client)
	Connectpool(client)
}

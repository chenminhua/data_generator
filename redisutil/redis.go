package redisutil

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/chenminhua/data_generator/constant"
	"github.com/chenminhua/data_generator/types"
	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client

func newRedisClient() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "R5wBXV3Djk", // no password set
		DB:       0,            // use default DB
	})
}

func CreateUser(start, end int) {
	ctx := context.TODO()
	msetcmd := []string{}
	for idx := start; idx < end; idx++ {
		msetcmd = append(msetcmd, fmt.Sprintf("u:%d", idx))
		u, _ := json.Marshal(types.NewUser(idx))
		msetcmd = append(msetcmd, string(u))
	}
	redisClient.MSet(ctx, msetcmd)
	//pipe := redisClient.Pipeline()
	//for idx := start; idx < end; idx++ {
	//	u, _ := json.Marshal(types.NewUser(idx))
	//	res, err := pipe.Set(ctx, fmt.Sprintf("u:%d", idx), string(u), 0).Result()
	//	print(res)
	//	if err != nil {
	//		println(err.Error())
	//	}
	//}
	//_, err := pipe.Exec(ctx)
	//if err != nil {
	//	println(err.Error())
	//}
}

func GenerateUsers() {
	ch := make(chan int)
	constant.LoadNamesData()
	newRedisClient()
	for i := 0; i < 32; i++ {
		go func() {
			for {
				idx := <-ch
				println(idx)
				CreateUser(idx, idx+100)
			}
		}()
	}
	for i := 0; i <= 1000000; i += 100 {
		println(i)
		ch <- i
	}
	close(ch)
}

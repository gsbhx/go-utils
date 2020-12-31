package main

import (
	"fmt"
	"github.com/gsbhx/go-utils"
	"github.com/gsbhx/go-utils/cache/redis"
	redisconf "github.com/gsbhx/go-utils/config/cache/redis"
	"github.com/gsbhx/go-utils/etcd"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main() {
	redisConf := redisconf.RedisConfigure{
		"127.0.0.1", 6379, 3000, 300, "hsuanyuen.cn",
	}
	redis.NewRedisClilent(redisConf)
	econf := clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	}
	err := etcd.EtcdInit(econf)
	if err != nil {
		fmt.Printf("%s", err.Error())
	}
	tools.InitTools()
	testRedis()
	testLog()
	testEtcd()
}

func testRedis() {

	status, err := tools.Hs.Cache.Redis.Set("aaa", 123995).String()
	fmt.Println(status, err)
	val, err := tools.Hs.Cache.Redis.Get("aaa").Int()
	fmt.Println(val, err)
}

func testLog() {
	err := tools.Hs.Error.ErrorWithCode.NewError(22522, "参数错误！")
	ce := tools.Hs.Error.ErrorWithCode.Assert(err).GetCode()
	fmt.Println(ce)
}

func testEtcd() {
	tools.Hs.Etcd.Put("fff", "222222222222323434")
	fmt.Println(tools.Hs.Etcd.Get("fff"))
}

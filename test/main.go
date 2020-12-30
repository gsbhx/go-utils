package main

import (
	"fmt"
	"git.hsuanyuen.cn/tools"
	"git.hsuanyuen.cn/tools/cache/redis"
	redisconf "git.hsuanyuen.cn/tools/config/cache/redis"
)

func main()  {

	testRedis()
	testLog()
}

func testRedis()  {
	redisConf:=redisconf.RedisConfigure{
		"127.0.0.1",6379,3000,300,"hsuanyuen.cn",
	}
	redis.NewRedisClilent(redisConf)
	var Hs=tools.Hs
	status,err:=Hs.Cache.Redis.Set("aaa",123995).String()
	fmt.Println(status,err)
	val,err:=Hs.Cache.Redis.Get("aaa").Int()
	fmt.Println(val,err)
}

func testLog()  {
	var Hs=tools.Hs
	err:=Hs.Error.ErrorWithCode.NewError(22522,"参数错误！");
	ce:= tools.Hs.Error.ErrorWithCode.Assert(err).GetCode()
	fmt.Println(ce)
}
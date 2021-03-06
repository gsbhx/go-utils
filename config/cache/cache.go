package cache

import redisconf "github.com/gsbhx/go-utils/config/cache/redis"
var CacheConf *CacheConfigure

type CacheConfigure struct {
	RedisConf *redisconf.RedisConfigure
}

func init()  {
	CacheConf=&CacheConfigure{
		RedisConf: redisconf.RedisConf,
	}
}
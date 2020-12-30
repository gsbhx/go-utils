package cache

import redisconf "tools/config/cache/redis"
var CacheConf *CacheConfigure

type CacheConfigure struct {
	RedisConf *redisconf.RedisConfigure
}

func init()  {
	CacheConf=&CacheConfigure{
		RedisConf: redisconf.RedisConf,
	}
}
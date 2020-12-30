package cache

import redisconf "git.hsuanyuen.cn/tools/config/cache/redis"
var CacheConf *CacheConfigure

type CacheConfigure struct {
	RedisConf *redisconf.RedisConfigure
}

func init()  {
	CacheConf=&CacheConfigure{
		RedisConf: redisconf.RedisConf,
	}
}
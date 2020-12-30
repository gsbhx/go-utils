package cache

import (
	hlog "tools/cache/log"
	"tools/cache/redis"
)

var HCache *Caches

type Caches struct {
	Redis redis.RedisInterface
	Log  *hlog.Logger
}

func init()  {
	HCache=new(Caches)
	HCache.Redis= new(redis.Redis)
	HCache.Log = new(hlog.Logger)
	HCache.Log.InitLogger("Logs")
}
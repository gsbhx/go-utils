package cache

import (
	hlog "github.com/gsbhx/go-utils/cache/log"
	"github.com/gsbhx/go-utils/cache/redis"
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
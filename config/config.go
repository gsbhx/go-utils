package config

import "github.com/gsbhx/go-utils/config/cache"

var Config *Conf
type Conf struct {
	Cache *cache.CacheConfigure
}

func init()  {
	Config=&Conf{
		Cache: cache.CacheConf,
	}
}
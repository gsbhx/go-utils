package config

import "github.com/gsbhx/tools/config/cache"

var Config *Conf
type Conf struct {
	Cache *cache.CacheConfigure
}

func init()  {
	Config=&Conf{
		Cache: cache.CacheConf,
	}
}
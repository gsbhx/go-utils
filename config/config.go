package config

import "git.hsuanyuen.cn/tools/config/cache"

var Config *Conf
type Conf struct {
	Cache *cache.CacheConfigure
}

func init()  {
	Config=&Conf{
		Cache: cache.CacheConf,
	}
}
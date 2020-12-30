package tools

import (
	"git.hsuanyuen.cn/tools/cache"
	"git.hsuanyuen.cn/tools/config"
	error2 "git.hsuanyuen.cn/tools/errors"
)

var Hs *hTools
var Hc *hConfig

type hTools struct {
	Cache *cache.Caches
	Error *error2.AllErrors
}

type hConfig struct {
	Config *config.Conf
}

func init() {
	Hs=new(hTools)
	Hc=new(hConfig)
	Hs.Cache = cache.HCache
	Hs.Error = new(error2.AllErrors)
	Hc.Config = config.Config
}

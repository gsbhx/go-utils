package tools

import (
	"github.com/gsbhx/go-utils/cache"
	"github.com/gsbhx/go-utils/config"
	error2 "github.com/gsbhx/go-utils/errors"
	"github.com/gsbhx/go-utils/etcd"
)

var Hs *hTools
var Hc *hConfig

type hTools struct {
	Cache *cache.Caches
	Error *error2.AllErrors
	Etcd *etcd.EtcdClient
}

type hConfig struct {
	Config *config.Conf
}

func InitTools() {
	//配置目录
	Hc=new(hConfig)
	Hc.Config = config.Config

	//工具目录
	Hs=new(hTools)
	Hs.Cache = cache.HCache
	Hs.Error = new(error2.AllErrors)
	Hs.Etcd= etcd.EClient

}

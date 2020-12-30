package etcd

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

var EClient *EtcdClient

type EtcdClient struct {
	config clientv3.Config
	client *clientv3.Client
}

func EtcdInit() (e *EtcdClient,err error) {
	e = new(EtcdClient)
	//初始化e.config TODO 将配置设置为通过ini文件提取
	e.config = clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	}
	if e.client,err= clientv3.New(e.config); err != nil {
		fmt.Println(err)
		return
	}
	return
}

func (e *EtcdClient) Get(name string) (s string,err error)  {
	kv := clientv3.NewKV(e.client)
	resq, err := kv.Get(context.TODO(), name);
	if err != nil {
		fmt.Println(err)
		return
	}
	s=string(resq.Kvs[0].Value)
	return

	return
}

func (e *EtcdClient)Put(name ,value string)(err error){
	kv := clientv3.NewKV(e.client)
	if _, err = kv.Put(context.TODO(), name, value); err != nil {
		fmt.Println(err)
		return
	}
	return
}


func (e *EtcdClient)new()(client *clientv3.Client,err error){
	if client, err = clientv3.New(e.config); err != nil {
		fmt.Println(err)
		return
	}
	return
}
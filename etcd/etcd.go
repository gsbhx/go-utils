package etcd

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
)

var EClient *EtcdClient

type EtcdClient struct {
	config clientv3.Config
	client *clientv3.Client
}

func EtcdInit(confs clientv3.Config) (err error) {
	EClient = new(EtcdClient)
	EClient.config = confs
	if EClient.client,err= clientv3.New(EClient.config); err != nil {
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
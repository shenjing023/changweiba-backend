package handler

import (
	"context"
	"fmt"
	"strings"
	"time"

	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc/resolver"
)

type Discovery struct {
	etcdCli     *clientv3.Client
	cc          resolver.ClientConn
	schema      string
	serviceName string
}

func NewDiscovery(conf clientv3.Config, schema, serviceName string) (resolver.Builder, error) {
	cli, err := clientv3.New(conf)
	if err != nil {
		return nil, err
	}
	return &Discovery{
		etcdCli:     cli,
		schema:      schema,
		serviceName: serviceName,
	}, nil
}

// Build 当调用`grpc.Dial()`时执行
func (d *Discovery) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	d.cc = cc
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// get key first
	prefix := GetPrefix(d.schema, d.serviceName)
	resp, err := d.etcdCli.Get(ctx, prefix, clientv3.WithPrefix())
	if err == nil {
		var addrList []resolver.Address
		for i := range resp.Kvs {
			addrList = append(addrList, resolver.Address{Addr: string(resp.Kvs[i].Value)})
		}
		d.cc.UpdateState(resolver.State{Addresses: addrList})
		go d.watch(prefix, addrList)
	} else {
		return nil, fmt.Errorf("etcd get failed, prefix: %s", prefix)
	}

	return d, err
}

func (d *Discovery) Scheme() string {
	return d.schema
}

func exists(addrList []resolver.Address, addr string) bool {
	for _, v := range addrList {
		if v.Addr == addr {
			return true
		}
	}
	return false
}

func remove(s []resolver.Address, addr string) ([]resolver.Address, bool) {
	for i := range s {
		if s[i].Addr == addr {
			s[i] = s[len(s)-1]
			return s[:len(s)-1], true
		}
	}
	return nil, false
}

func (d *Discovery) watch(prefix string, addrList []resolver.Address) {
	rch := d.etcdCli.Watch(context.Background(), prefix, clientv3.WithPrefix(), clientv3.WithPrefix())
	for n := range rch {
		flag := 0
		for _, ev := range n.Events {
			switch ev.Type {
			case mvccpb.PUT:
				if !exists(addrList, string(ev.Kv.Value)) {
					flag = 1
					addrList = append(addrList, resolver.Address{Addr: string(ev.Kv.Value)})
					fmt.Println("after add, new list: ", addrList)
				}
			case mvccpb.DELETE:
				fmt.Println("remove addr key: ", string(ev.Kv.Key), "value:", string(ev.Kv.Value))
				i := strings.LastIndexAny(string(ev.Kv.Key), "/")
				if i < 0 {
					return
				}
				t := string(ev.Kv.Key)[i+1:]
				fmt.Println("remove addr key: ", string(ev.Kv.Key), "value:", string(ev.Kv.Value), "addr:", t)
				if s, ok := remove(addrList, t); ok {
					flag = 1
					addrList = s
					fmt.Println("after remove, new list: ", addrList)
				}
			}
		}

		if flag == 1 {
			d.cc.UpdateState(resolver.State{Addresses: addrList})
			fmt.Println("update: ", addrList)
		}
	}
}

func (d *Discovery) ResolveNow(rn resolver.ResolveNowOptions) {
}

// Close 当调用`grpc.ClientConn.Close()`时执行
func (d *Discovery) Close() {
	d.etcdCli.Close()
}

// "%s:///%s/"
func GetPrefix(schema, serviceName string) string {
	return fmt.Sprintf("%s:///%s/", schema, serviceName)
}

package main

import (
	"context"
	"fmt"
	"net"

	clientv3 "go.etcd.io/etcd/client/v3"
)

type Register struct {
	etcdCli *clientv3.Client
	ttl     int64
	key     string
}

type Option func(*Register)

func WithTTL(ttl int64) Option {
	return func(r *Register) {
		r.ttl = ttl
	}
}

func NewRegister(conf clientv3.Config, schema, serviceName, host, port string, opts ...Option) (*Register, error) {
	cli, err := clientv3.New(conf)
	if err != nil {
		return nil, fmt.Errorf("create etcd clientv3 client failed, errmsg:%v", err)
	}

	r := &Register{
		etcdCli: cli,
		ttl:     10,
	}
	for _, o := range opts {
		o(r)
	}

	//lease
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	resp, err := cli.Grant(ctx, r.ttl)
	if err != nil {
		return nil, fmt.Errorf("grant failed, errmsg:%v", err)
	}

	//  schema:///serviceName/ip:port ->ip:port
	serviceValue := net.JoinHostPort(host, port)
	serviceKey := GetPrefix(schema, serviceName) + serviceValue
	r.key = serviceKey

	//set key->value
	if _, err := cli.Put(ctx, serviceKey, serviceValue, clientv3.WithLease(resp.ID)); err != nil {
		return nil, fmt.Errorf("put failed, errmsg:%v， key:%s, value:%s", err, serviceKey, serviceValue)
	}

	//keepalive
	kresp, err := cli.KeepAlive(context.Background(), resp.ID)
	if err != nil {
		return nil, fmt.Errorf("keepalive faild, errmsg:%v, lease id:%d", err, resp.ID)
	}

	go func() {
	FLOOP:
		for v := range kresp {
			if v == nil {
				fmt.Println("etcd keepalive closed")
				break FLOOP
			}
		}
	}()

	return r, nil
}

// Close 注销服务
func (r *Register) Close() error {
	if _, err := r.etcdCli.Delete(context.Background(), r.key); err != nil {
		return err
	}
	return r.etcdCli.Close()
}

// "%s:///%s/"
func GetPrefix(schema, serviceName string) string {
	return fmt.Sprintf("%s:///%s/", schema, serviceName)
}

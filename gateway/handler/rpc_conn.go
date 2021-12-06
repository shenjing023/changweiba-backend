package handler

import (
	"context"
	"errors"
	"fmt"
	"time"

	"gateway/common"
	"gateway/conf"

	log "github.com/shenjing023/llog"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/resolver"
)

//rpc连接
var (
	AccountConn *grpc.ClientConn
	PostConn    *grpc.ClientConn
)

const serviceSchema = "svc"

// InitGRPCConn init grpc conn
func InitGRPCConn() {
	var err error
	etcdConf := clientv3.Config{
		Endpoints:   []string{fmt.Sprintf("%s:%d", conf.Cfg.Etcd.Host, conf.Cfg.Etcd.Port)},
		DialTimeout: time.Second * 5,
	}
	account, err := NewDiscovery(etcdConf, "svc-"+conf.Cfg.AccountSvcName, conf.Cfg.AccountSvcName)
	if err != nil {
		panic(err)
	}
	resolver.Register(account)

	post, err := NewDiscovery(etcdConf, "svc-"+conf.Cfg.PostSvcName, conf.Cfg.PostSvcName)
	if err != nil {
		panic(err)
	}
	resolver.Register(post)

	AccountConn, err = grpc.DialContext(context.Background(),
		GetPrefix("svc-"+conf.Cfg.AccountSvcName, conf.Cfg.AccountSvcName),
		grpc.WithDefaultServiceConfig(`{"LoadBalancingPolicy": "round_robin"}`),
		grpc.WithInsecure(),
		// grpc.WithUnaryInterceptor(unaryHeaderInterceptor),
	)
	if err != nil {
		log.Fatalf("fail to accountRPC dial: %+v", err)
	}

	PostConn, err = grpc.DialContext(context.Background(),
		GetPrefix("svc-"+conf.Cfg.PostSvcName, conf.Cfg.PostSvcName),
		grpc.WithDefaultServiceConfig(`{"LoadBalancingPolicy": "round_robin"}`),
		grpc.WithInsecure(),
		// grpc.WithUnaryInterceptor(unaryHeaderInterceptor),
	)
	if err != nil {
		log.Fatalf("fail to postsRPC dial: %+v", err)
	}
}

// StopGRPCConn 关闭rpc连接
func StopGRPCConn() {
	if AccountConn != nil {
		AccountConn.Close()
	}
	if PostConn != nil {
		PostConn.Close()
	}
}

// unaryHeaderInterceptor get http header to metadata for opentracing
func unaryHeaderInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	gc, err := common.GinContextFromContext(ctx)
	if err != nil {
		log.Errorf("%+v", err)
		return errors.New(ServiceError)
	}
	var (
		otHeaders = []string{
			"X-Request-Id",
			"X-B3-Parentspanid",
			"X-B3-Sampled",
			"X-B3-Spanid",
			"X-B3-Traceid",
		}
		pairs []string
	)

	for _, h := range otHeaders {
		if v := gc.Request.Header.Get(h); len(v) > 0 {
			pairs = append(pairs, h, v)
		}
	}
	header := metadata.Pairs(pairs...)
	ctx = metadata.NewOutgoingContext(ctx, header)
	return invoker(ctx, method, req, reply, cc, opts...)
}

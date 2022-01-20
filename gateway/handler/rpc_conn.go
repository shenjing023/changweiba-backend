package handler

import (
	"context"
	"fmt"
	"time"

	"gateway/conf"

	pb "gateway/pb"

	log "github.com/shenjing023/llog"
	vp_client "github.com/shenjing023/vivy-polaris/client"
	"github.com/shenjing023/vivy-polaris/contrib/registry"
	"github.com/shenjing023/vivy-polaris/contrib/tracing"
	"github.com/shenjing023/vivy-polaris/options"
	clientv3 "go.etcd.io/etcd/client/v3"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"google.golang.org/grpc"
)

//rpc连接
var (
	AccountConn *grpc.ClientConn
	PostConn    *grpc.ClientConn
	StockConn   *grpc.ClientConn
	tp          *sdktrace.TracerProvider
)

// InitGRPCConn init grpc conn
func InitGRPCConn() {
	var err error
	etcdConf := clientv3.Config{
		Endpoints:   []string{fmt.Sprintf("%s:%d", conf.Cfg.Etcd.Host, conf.Cfg.Etcd.Port)},
		DialTimeout: time.Second * 5,
	}

	// jeager tracing
	tp, err = tracing.NewJaegerTracerProvider(conf.Cfg.JaegerCollectURL, "gateway-client")
	if err != nil {
		log.Fatalf("new JaegerTracerProvider error: %+v", err)
	}

	AccountConn, err = vp_client.NewClientConn(registry.GetServiceTarget(pb.AccountService_ServiceDesc), options.WithEtcdDiscovery(etcdConf, pb.AccountService_ServiceDesc),
		options.WithInsecure(), options.WithRRLB(), options.WithClientTracing(tp))
	if err != nil {
		log.Fatalf("fail to accountRPC dial: %+v", err)
	}

	PostConn, err = vp_client.NewClientConn(registry.GetServiceTarget(pb.PostService_ServiceDesc), options.WithEtcdDiscovery(etcdConf, pb.PostService_ServiceDesc),
		options.WithInsecure(), options.WithRRLB(), options.WithClientTracing(tp))
	if err != nil {
		log.Fatalf("fail to postsRPC dial: %+v", err)
	}

	StockConn, err = vp_client.NewClientConn(registry.GetServiceTarget(pb.StockService_ServiceDesc), options.WithEtcdDiscovery(etcdConf, pb.StockService_ServiceDesc),
		options.WithInsecure(), options.WithRRLB(), options.WithClientTracing(tp))
	if err != nil {
		log.Fatalf("fail to stockRPC dial: %+v", err)
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
	if StockConn != nil {
		StockConn.Close()
	}
}

func StopTracer() {
	tp.Shutdown(context.Background())
}

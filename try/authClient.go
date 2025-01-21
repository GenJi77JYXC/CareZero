package main

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc/metadata"
	"www.genji.xin/backend/CareZero/authServer/auth"
)

func main() {
	NewAuthClient()
}

func NewAuthClient() {
	clientConf := zrpc.RpcClientConf{}
	conf.FillDefault(&clientConf)
	clientConf.Etcd = discov.EtcdConf{
		Hosts: []string{"127.0.0.1:2379"},
		Key:   "authorization.rpc",
	}
	conn := zrpc.MustNewClient(clientConf)
	client := auth.NewAuthServiceClient(conn.Conn())
	md := metadata.New(map[string]string{
		"authorization": fmt.Sprintf("Bearer %s", clientConf.Target),
	})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	resp, err := client.Ping(ctx, &auth.Request{})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)

}

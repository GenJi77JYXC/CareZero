package middleware

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
)

func AuthInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler, // 业务处理函数
) (interface{}, error) {
	// 请求前的逻辑
	log.Printf("请求前: 接收到请求, 方法: %s, 请求参数: %v", info.FullMethod, req)
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Printf("获取上下文错误：客户端没有传...")
		return nil, status.Error(codes.Unauthenticated, "缺少 Metadata")
	}
	token := md["authorization"]
	if len(token) == 0 {
		log.Printf("没有从上下文中拿到token")
	}

	// jwt认证

	// 调用实际的业务处理逻辑（进入服务端的 Handler）
	resp, err := handler(ctx, req)

	// 请求后的逻辑
	if err != nil {
		log.Printf("请求后: 方法: %s, 处理发生错误: %v", info.FullMethod, err)
	} else {
		log.Printf("请求后: 方法: %s, 响应参数: %v", info.FullMethod, resp)
	}

	// 返回响应和错误
	return resp, err
}

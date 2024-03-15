package api

import (
	"context"
	"github.com/lightningnetwork/lnd/lnrpc/invoicesrpc"
	"google.golang.org/grpc"
	"log"
)

func AddHoldInvoice() *invoicesrpc.AddHoldInvoiceResp {
	// 读取参数
	grpcHost := getEnv("RPC_SERVER")
	tlsCertPath := getEnv("TLS_CERT_PATH")
	macaroonPath := getEnv("MACAROON_PATH")
	// 处理证书和macaroon
	creds := newTlsCert(tlsCertPath)
	macaroon := getMacaroon(macaroonPath)
	// 连接到grpc服务器
	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(newMacaroonCredential(macaroon)))
	if err != nil {
		log.Fatalf("did not connect: grpc.Dial: %v", err)
	}
	// 匿名函数延迟关闭grpc连接
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("conn Close Error: %v", err)
		}
	}(conn)
	// 创建客户端
	client := invoicesrpc.NewInvoicesClient(conn)
	// 构建请求
	request := &invoicesrpc.AddHoldInvoiceRequest{}
	// 得到响应
	response, err := client.AddHoldInvoice(context.Background(), request)
	if err != nil {
		log.Fatalf("invoicesrpc AddHoldInvoice Error: %v", err)
	}
	// 处理结果
	return response
}

package api

import (
	"context"
	"github.com/lightningnetwork/lnd/lnrpc"
	"google.golang.org/grpc"
	"log"
)

func AbandonChannel() {}

func AddInvoice() {}

func BakeMacaroon() {}

func BatchOpenChannel() {}

func ChannelAcceptor() {}

func ChannelBalance() {}

func CheckMacaroonPermissions() {}

func CloseChannel() {}

func ClosedChannels() {}

func ConnectPeer() {}

func DebugLevel() {}

func DecodePayReq() {}

func DeleteAllPayments() {}

func DeleteMacaroonID() {}

func DeletePayment() {}

func DescribeGraph() {}

func DisconnectPeer() {}

func EstimateFee() {}

func ExportAllChannelBackups() {}

func ExportChannelBackup() {}

func FeeReport() {}

func ForwardingHistory() {}

func FundingStateStep() {}

func GetChanInfo() {}

func GetDebugInfo() {}

func GetInfo() *lnrpc.GetInfoResponse {
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
	client := lnrpc.NewLightningClient(conn)
	// 构建请求
	request := &lnrpc.GetInfoRequest{}
	// 得到响应
	response, err := client.GetInfo(context.Background(), request)
	if err != nil {
		log.Fatalf("lnrpc GetInfo Error: %v", err)
	}
	// 处理结果
	return response
}

func GetNetworkInfo() {}

func GetNodeInfo() {}

func GetNodeMetrics() {}

func GetRecoveryInfo() {}

func GetTransactions() {}

func ListAliases() {}

func ListChannels() {}

func ListInvoices() {}

func ListMacaroonIDs() {}

func ListPayments() {}

func ListPeers() {}

func ListPermissions() {}

func ListUnspent() {}

func LookupHtlcResolution() {}

func LookupInvoice() {}

func NewAddress() {}

func OpenChannelSync() {}

func OpenChannel() *lnrpc.Lightning_OpenChannelClient {
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
	client := lnrpc.NewLightningClient(conn)
	// 构建请求
	request := &lnrpc.OpenChannelRequest{
		SatPerVbyte:                0,
		NodePubkey:                 nil,
		NodePubkeyString:           "",
		LocalFundingAmount:         0,
		PushSat:                    0,
		TargetConf:                 0,
		SatPerByte:                 0,
		Private:                    false,
		MinHtlcMsat:                0,
		RemoteCsvDelay:             0,
		MinConfs:                   0,
		SpendUnconfirmed:           false,
		CloseAddress:               "",
		FundingShim:                nil,
		RemoteMaxValueInFlightMsat: 0,
		RemoteMaxHtlcs:             0,
		MaxLocalCsv:                0,
		CommitmentType:             0,
		ZeroConf:                   false,
		ScidAlias:                  false,
		BaseFee:                    0,
		FeeRate:                    0,
		UseBaseFee:                 false,
		UseFeeRate:                 false,
		RemoteChanReserveSat:       0,
		FundMax:                    false,
		Memo:                       "",
		Outpoints:                  nil,
	}
	// 得到响应
	response, err := client.OpenChannel(context.Background(), request)
	if err != nil {
		log.Fatalf("lnrpc OpenChannel Error: %v", err)
	}
	// 处理结果
	return &response
}

func PendingChannels() {}

func QueryRoutes() {}

func RegisterRPCMiddleware() {}

func RestoreChannelBackups() {}

func SendCoins() {}

func SendCustomMessage() {}

func SendMany() {}

func SendPaymentSync() {}

func SendPayment() {}

func SendToRouteSync() {}

func SendToRoute() {}

func SignMessage() {}

func StopDaemon() {}

func SubscribeChannelBackups() {}

func SubscribeChannelEvents() {}

func SubscribeChannelGraph() {}

func SubscribeCustomMessages() {}

func SubscribeInvoices() {}

func SubscribePeerEvents() {}

func SubscribeTransactions() {}

func UpdateChannelPolicy() {}

func VerifyChanBackup() {}

func VerifyMessage() {}

func WalletBalance() {}

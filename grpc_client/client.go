package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc_s/grpc_client/auth"
	"grpc_s/services"
	"log"
	"strconv"
)

func main() {
	// grpc.WithTransportCredentials() 没有证书报错
	rpcAuth := &auth.RpcAuth{Username: "admin", Pass: "admin"}
	c, err := grpc.Dial(":50081", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithPerRPCCredentials(rpcAuth))
	if err != nil {
		log.Fatal("服务端连接出错", err)
	}
	defer c.Close()

	cs := services.NewSearchServiceClient(c)
	//rs, err := cs.SearchUser(context.Background(), &services.SearchParam{Username: "cxl98655adfads"})
	//if err != nil {
	//	log.Fatal("有错误", err.Error())
	//}
	steam, _ := cs.GetUserSteam(context.Background())
	for i := 0; i < 10; i++ {
		_ = steam.Send(&services.UserRequest{Username: "user_" + strconv.Itoa(i+1)})
	}
	recv, err := steam.CloseAndRecv()
	log.Println(recv.GetUsers())
}

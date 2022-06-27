package main

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"grpc_s/services"
	"log"
	"net"
)

func main() {
	var authInterceptor grpc.UnaryServerInterceptor //UnaryServerInterceptor
	authInterceptor = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		err = RpcAuth(ctx)
		if err != nil {
			return nil, err
		}
		return handler(ctx, req)
	}
	grpcService := grpc.NewServer(grpc.UnaryInterceptor(authInterceptor))
	services.RegisterSearchServiceServer(grpcService, new(services.SearchService))
	n, _ := net.Listen("tcp", ":50081")
	err := grpcService.Serve(n)
	log.Println(err)
}

func RpcAuth(ctx context.Context) error {
	if md, ok := metadata.FromIncomingContext(ctx); !ok {
		return errors.New("params error")
	} else {
		var (
			user string
			pass string
		)
		if users, exists := md["user"]; exists {
			user = users[0]
		} else {
			return errors.New("params user error")
		}
		if passes, exists := md["pass"]; exists {
			pass = passes[0]
		} else {
			return errors.New("params pass error")
		}
		if user != "admin" || pass != "admin" {
			return errors.New("auth error")
		}
		return nil
	}
}

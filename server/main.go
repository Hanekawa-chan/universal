package main

import (
	"context"
	"github.com/Hanekawa-chan/universal/protoc"

	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
)

type UniversalServer struct {
	protoc.UnimplementedUniversalServer
	mu          sync.Mutex
	authed      map[string]bool
	credentials map[string]string
	connections map[string]chan *protoc.Message
}

func newServer() *UniversalServer {
	var s = &UniversalServer{
		authed:      make(map[string]bool),
		credentials: make(map[string]string),
		connections: make(map[string]chan *protoc.Message),
	}
	return s
}

func (u *UniversalServer) SignUp(ctx context.Context, user *protoc.User) (*protoc.Response, error) {
	for i := range u.credentials {
		if i == user.Name {
			return &protoc.Response{
				Response: "ALREADY EXISTS",
			}, nil
		}
	}
	u.credentials[user.Name] = user.Password
	return &protoc.Response{
		Response: "OK",
	}, nil
}

func (u *UniversalServer) SignIn(ctx context.Context, user *protoc.User) (*protoc.Response, error) {
	for i, pass := range u.credentials {
		if i == user.Name {
			if pass == user.Password {
				u.authed[i] = true
				return &protoc.Response{
					Response: "OK",
				}, nil
			} else {
				u.authed[i] = false
				return &protoc.Response{
					Response: "WRONG PASSWORD",
				}, nil
			}
		}
	}
	return &protoc.Response{
		Response: "CAN'T FIND USER",
	}, nil
}

func (u *UniversalServer) SignOut(ctx context.Context, user *protoc.User) (*protoc.Response, error) {
	for i := range u.credentials {
		if i == user.Name {
			u.authed[i] = false
			return &protoc.Response{
				Response: "OK",
			}, nil
		}
	}
	return &protoc.Response{
		Response: "CAN'T FIND USER",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:5565")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	protoc.RegisterUniversalServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}

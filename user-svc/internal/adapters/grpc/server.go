package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/nico-phil/notification-proto/golang/user"
	"github.com/nico-phil/notification/user-svc/config"
	"github.com/nico-phil/notification/user-svc/internal/ports"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Adapter struct {
	api ports.APIPort
	port int
	server *grpc.Server
	user.UnimplementedUserServer
}

func NewAdapter(api ports.APIPort, port int) *Adapter {
	return &Adapter {
		api: api,
		port: port,
	}
}

func(a *Adapter) Run(){
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatal("failed to create listener")
	}

	grpcServer := grpc.NewServer()
	a.server = grpcServer
	user.RegisterUserServer(grpcServer, a)

	if config.GetEnv() == "development"{
		reflection.Register(grpcServer)
	}

	err = grpcServer.Serve(listen) 
	if err != nil {
		log.Fatalf("failed to serve grpc on port %d", a.port)
	}


}
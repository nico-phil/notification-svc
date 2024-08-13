package grpc

import (
	"fmt"
	"log"
	"net"

	notifs "github.com/nico-phil/notification-proto/golang/notification"
	"github.com/nico-phil/notification/config"
	"github.com/nico-phil/notification/internal/ports"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)


type Adapter struct {
	api ports.APIPort
	server *grpc.Server
	port int
	notifs.UnimplementedNotificationServer
}

func NewAdapter(api ports.APIPort, port int) *Adapter {
	return &Adapter{
		port: port,
		api:api,
	}
}

func(a Adapter) Run(){

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen to port %d, err: %v", a.port, err)
	}

	grpcServer := grpc.NewServer()
	a.server = grpcServer
	notifs.RegisterNotificationServer(grpcServer, a)

	if config.GetEnv() == "development" {
		reflection.Register(grpcServer)
	}

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve grpc on port:%d", a.port)
	}
	
}

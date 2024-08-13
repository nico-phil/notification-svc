package main

import (
	"github.com/nico-phil/notification/internal/adapters/grpc"
	"github.com/nico-phil/notification/internal/application/core/api"
)

func main(){

	application := api.NewApplication()

	grpcAdapter := grpc.NewAdapter(application, 3000)
	grpcAdapter.Run()

}
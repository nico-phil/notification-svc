package main

import (
	"log"

	"github.com/nico-phil/notification/user-svc/config"
	"github.com/nico-phil/notification/user-svc/internal/adapters/db"
	"github.com/nico-phil/notification/user-svc/internal/adapters/grpc"
	"github.com/nico-phil/notification/user-svc/internal/application/core/api"
)

func main(){

	dbAdapter, err := db.NewAdapter(config.GetDSN())
	if err != nil {
		log.Fatalf("failed to connect to postgres %v", err)
	}
	log.Println("sucessfully connect to posgres")
	application := api.NewApplication(dbAdapter)

	grpcAdapter := grpc.NewAdapter(application, config.GetAppPort())

	grpcAdapter.Run()
}
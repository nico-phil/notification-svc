package main

import (
	"log"

	"github.com/nico-phil/notification/config"
	"github.com/nico-phil/notification/internal/adapters/db"
	"github.com/nico-phil/notification/internal/adapters/grpc"
	"github.com/nico-phil/notification/internal/adapters/producer"
	"github.com/nico-phil/notification/internal/application/core/api"
)

func main(){
	dbAdapter, err := db.NewAdapter(config.GetDbDSN())
	if err != nil {
		log.Fatal("cannot connect to db...")
	}
	log.Println("succesffuly connect to postgres")

	producerAdapter, err := producer.NewAdapter([]string{"localhost:9092"})
	if err != nil {
		log.Fatalf("failed to connect to kafka err: %v", err)
	}
	
	application := api.NewApplication(producerAdapter, dbAdapter)

	grpcAdapter := grpc.NewAdapter(application, 3000)
	grpcAdapter.Run()

}
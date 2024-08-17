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

	producerAdapter, err := producer.NewAdapter([]string{"localhost:9092"})
	if err != nil {
		log.Fatalf("failed to connect to kafka err: %v", err)
	}
	dbAdapter := &db.DBModel{Devices: []db.DeviceEntity{}}
	_, err  = db.NewAdapter(config.GetDbDSN())
	if err != nil {
		log.Fatal("cannot connect to db...")
	}
	log.Println("app succesffuly connect postgre")
	application := api.NewApplication(producerAdapter, dbAdapter)

	grpcAdapter := grpc.NewAdapter(application, 3000)
	grpcAdapter.Run()

}
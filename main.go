package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/micro/go-micro"
	pb "github.com/oentoro/consignment-service/proto/consignment"
	vesselProto "github.com/oentoro/vessel-service/proto/vessel"
)

// Import the generated protobuf code

const (
	port        = ":50051"
	defaultHost = "datastore:27017"
)

func main() {

	srv := micro.NewService(micro.Name("consignment.service"))

	srv.Init()

	uri := os.Getenv("DB_HOST")

	if uri == "" {
		uri = defaultHost
	}

	client, err := CreateClient(uri)
	if err != nil {
		log.Panic(err)
	}

	defer client.Disconnect(context.TODO())

	consignmentCollection := client.Database("shippy").Collection("consignments")

	repository := &MongoRepository{consignmentCollection}
	vesselClient := vesselProto.NewVesselServiceClient("vessel.service", srv.Client())
	h := &handler{repository, vesselClient}

	pb.RegisterConsignmentServiceHandler(srv.Server(), h)

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}

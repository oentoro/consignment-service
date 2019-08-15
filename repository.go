package main

import (
	"context"

	pb "github.com/oentoro/consignment-service/proto/consignment"
	"go.mongodb.org/mongo-driver/mongo"
)

//Repository - consignment's repos
type Repository interface {
	Create(consignment *pb.Consignment) error
	GetAll() ([]*pb.Consignment, error)
}

//MongoRepository - collection from mongo database
type MongoRepository struct {
	collection *mongo.Collection
}

//Create - create new consignment
func (repository *MongoRepository) Create(consignment *pb.Consignment) error {
	_, err := repository.collection.InsertOne(context.Background(), consignment)
	return err
}

//GetAll - get all consignment from repos
func (repository *MongoRepository) GetAll() ([]*pb.Consignment, error) {
	cur, err := repository.collection.Find(context.Background(), nil, nil)
	var consignments []*pb.Consignment
	for cur.Next(context.Background()) {
		var consignment *pb.Consignment
		if err := cur.Decode(&consignment); err != nil {
			return nil, err
		}
		consignments = append(consignments, consignment)
	}
	return consignments, err
}

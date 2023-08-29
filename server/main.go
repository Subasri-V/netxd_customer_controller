package main

import (
	"context"
	"fmt"
	"net"
	"netxd_customer_controller/config"
	"netxd_customer_controller/constants"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func initDatabase(client *mongo.Client) {
	customerCollection := config.GetCollection(client, constants.DatabaseName, "customers")

}
func main() {
	mongoClient, err := config.ConnectDataBase()
	defer mongoClient.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
	initDatabase(mongoClient)

	lis, err := net.Listen("tcp", constants.Port)
	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()
}

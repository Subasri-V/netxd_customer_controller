package main

import (
	pro "github.com/Subasri-V/application-new/netxd_customer/netxd"
	"github.com/Subasri-V/application-new/netxd_customer_dal/services"
	"context"
	"fmt"
	"net"
	"github.com/Subasri-V/application-new/netxd_customer_controller/config"
	"github.com/Subasri-V/application-new/netxd_customer_controller/constants"
	controllers "github.com/Subasri-V/application-new/netxd_customer_controller/controller"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func initDatabase(client *mongo.Client) {
	customerCollection := config.GetCollection(client, constants.DatabaseName, "customers")
	pro.CustomerDetails = services.InitializeCustomerService(customerCollection, context.Background())
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

	pro.RegisterCustomerDetailsServer(s, &controllers.RPCServer{})

	fmt.Println("Server listening on", constants.Port)
	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %v", err)
	}
}

package main

import (
	"context"
	"fmt"
	"net"

	cus "github.com/Subasri-V/application-new/netxd_customer/netxd"

	"github.com/Subasri-V/application-new/netxd_customer_controller/config"
	"github.com/Subasri-V/application-new/netxd_customer_controller/constants"
	controllers "github.com/Subasri-V/application-new/netxd_customer_controller/controller"
	"github.com/Subasri-V/application-new/netxd_customer_dal/services"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func initDatabase(client *mongo.Client) {
	customerCollection := config.GetCollection(client, constants.DatabaseName, "customers")
	controllers.CustomerDetails = services.InitializeCustomerService(context.Background(), customerCollection)
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

	cus.RegisterCustomerDetailsServer(s, &controllers.RPCServer{})

	fmt.Println("Server listening on", constants.Port)
	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %v", err)
	}
}

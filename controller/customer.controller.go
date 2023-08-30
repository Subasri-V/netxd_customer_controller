package controller

import (
	"context"

	cus "github.com/Subasri-V/application-new/netxd_customer/netxd"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Subasri-V/application-new/netxd_customer_dal/interfaces"
	"github.com/Subasri-V/application-new/netxd_customer_dal/models"
)

type RPCServer struct {
	cus.UnimplementedCustomerDetailsServer
}

var (
	customerCollection *mongo.Collection
	CustomerDetails interfaces.ICustomer
)

func (s*RPCServer) CreateCustomer(ctx context.Context,req *cus.CustomerRequest)(*cus.CustomerResponse,error){
	newCust:=&models.CustomerDetails{Customerid: req.Customerid,Firstname: req.Firstname,Lastname: req.Lastname,Bankid: req.Bankid,Balance: float64(req.Balance),IsActive: req.IsActive}
	result,err:=CustomerDetails.CreateCustomer(newCust)

	if err!=nil{
		return nil,err
	} else {
		responseCustomer:=&cus.CustomerResponse{
			Customerid: result.Customerid,
	
		}
		return responseCustomer,nil
	}
}
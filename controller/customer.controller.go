package controller

import (
	"context"

	cus "github.com/Subasri-V/application-new/netxd_customer/netxd"

	"github.com/Subasri-V/application-new/netxd_customer_dal/interfaces"
	"github.com/Subasri-V/application-new/netxd_customer_dal/models"
)

type RPCServer struct {
	cus.UnimplementedCustomerDetailsServer
}

var (
	//customerCollection *mongo.Collection
	CustomerDetails interfaces.ICustomer
)

func (s *RPCServer) CreateCustomer(ctx context.Context, req *cus.CustomerRequest) (*cus.CustomerResponse, error) {
	newCust := &models.CustomerDetails{Customerid: req.Customerid, Firstname: req.Firstname, Lastname: req.Lastname, Bankid: req.Bankid, Balance: int32(req.Balance), IsActive: req.IsActive}
	result, err := CustomerDetails.CreateCustomer(newCust)

	if err != nil {
		return nil, err
	} else {
		responseCustomer := &cus.CustomerResponse{
			Customerid: result.Customerid,
		}
		return responseCustomer, nil
	}
}

func (s*RPCServer) GetCustomerById(ctx context.Context,req *cus.IdReq)(*cus.IdRes,error){
	//var getCust int32
	getCust:=req.Customerid

	result,err:=CustomerDetails.GetCustomerById(getCust)

	if err != nil {
		return nil, err
	} else {
		responseCustomer := &cus.IdRes{
			Customerid: result.Customerid,
			Firstname: result.Firstname,
			Balance: result.Balance,
		}
		return responseCustomer, nil
	}
}

func(s* RPCServer) DeleteCustomerById(ctx context.Context,req *cus.DeleteReq)(*cus.DeleteRes,error){
	deleteCust:=req.Customerid

	_,err:=CustomerDetails.DeleteCustomerById(deleteCust)
	if err != nil {
		return nil, err
	}
	return &cus.DeleteRes{
		Message: "success",
	},nil

}

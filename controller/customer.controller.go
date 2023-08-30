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
	CustomerDetails interfaces.ICustomer
)

func (s *RPCServer) CreateCustomer(ctx context.Context, req *cus.customerRequest) (*cus.customerResponse, error) {
	newCust := &models.CustomerDetails{Firstname: req.Firstname}
	res, err := CustomerDetails.CreateCustomer(newCust)
	if err != nil {
		return nil, err
	} else {
		resCust := &cus.customerResponse{
			Firstname: res.Firstname,
		}
		return resCust, nil
	}
}

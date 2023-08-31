package controller

import (
	"context"

	"github.com/SinekaV/grpc-dal/interfaces"
	"github.com/SinekaV/grpc-dal/models"

	c "github.com/SinekaV/grpc-proto/customer"
)
type RPCServer struct{
	c.UnimplementedCustomerServiceServer
}

var(
	CustomerService interfaces.ICustomer
)

func(s *RPCServer)CreateCustomer(ctx context.Context,req * c.CustomerRequest)(*c.CustomerResponse,error){
	dbProfile:=&models.CustomerRequest{
		CustomerId: req.CustomerId,
		Firstname:  req.Firstname,
		Lastname:   req.Lastname,
		BankId:     req.BankId,
		Balance:    req.Balance,
	}
	res,err:=CustomerService.CreateCustomer(dbProfile)
	if err != nil {
		return nil, err
	}else {
		responseProfile := &c.CustomerResponse{
			CustomerId: res.CustomerId,
			CreatedAt: res.CreatedAt.Format("2006-01-02 15:04:05"),
		}
		return responseProfile, nil
	}

}
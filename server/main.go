package main

import (
	"context"
	"fmt"
	"net"

	"github.com/SinekaV/grpc-config/config"
	"github.com/SinekaV/grpc-connectors/controller"
	"github.com/SinekaV/grpc-constants/constants"
	"github.com/SinekaV/grpc-dal/services"
	c "github.com/SinekaV/grpc-proto/customer"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func intiDatabase(client *mongo.Client){
	customerCollection:=config.GetCollection(client,"BankDatabase","Customer")
	controller.CustomerService=services.InitCustomerService(customerCollection,context.Background())
}

func main(){
	mongoclient,err:=config.ConnectDatabase()
	defer mongoclient.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
	intiDatabase(mongoclient)
	lis,err:=net.Listen("tcp", constants.Port)//capture the code..gaining the access to a port"net.Listen"
	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
		return
	}
	s:=grpc.NewServer()//it will give the access to the memory ocation where we can map our code..creating the instance of grpc server
	c.RegisterCustomerServiceServer(s,&controller.RPCServer{})

	fmt.Println("sever listening on",constants.Port)
	if err := s.Serve(lis); err != nil {//grpc server.serve on the port
		fmt.Printf("Failed to serve: %v", err)
	}
}
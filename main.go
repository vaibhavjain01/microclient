package main

import (
	"context"
	"fmt"
	"time"

	"github.com/micro/micro/v3/service"
	pb "github.com/vaibhavjain01/microclient/proto"
)

func main() {
	// create and initialise a new service
	srv := service.New()

	// create the proto client for helloworld
	client := pb.NewHelloService("hello", srv.Client())

	// call an endpoint on the service
	rsp, err := client.Call(context.Background(), &pb.Request{
		Name: "John",
	})
	if err != nil {
		fmt.Println("Error calling service_one: ", err)
		return
	}

	// print the response
	fmt.Println("Response: ", rsp.Msg)
	
	// let's delay the process for exiting for reasons you'll see below
	time.Sleep(time.Second * 5)
}
package main

import (
	"github.com/micro/go-micro"
	pb "github.com/johynpapin/yonjuuni/email-service/proto/email"
	"fmt"
)

func main() {
	// Create a new service
	srv := micro.NewService(
		micro.Name("jp.yonjuuni.srv.email"),
		micro.Version("latest"),
	)

	// Register a new Subscriber
	micro.RegisterSubscriber("user.created", srv.Server(), new(Subscriber))

	// Init will parse the command line flags
	srv.Init()

	// Register handler
	pb.RegisterEmailServiceHandler(srv.Server(), &Service{})

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}

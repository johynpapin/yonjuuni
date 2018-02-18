package main

import (
	"fmt"
	pb "github.com/johynpapin/yonjuuni/user-service/proto/user"
	"github.com/micro/go-micro"
	"log"
)

func main() {
	// Connect to the database
	db, err := CreateConnection()
	if err != nil {
		log.Panicf("could not connect to database: %v", err)
	}
	defer db.Close()

	db.AutoMigrate(pb.User{})

	repo := &UserRepository{db}

	tokenService := &TokenService{repo}

	// Create a new service
	srv := micro.NewService(
		micro.Name("jp.yonjuuni.srv.user"),
		micro.Version("latest"),
	)

	// Create a new publisher
	publisher := micro.NewPublisher("user.created", srv.Client())

	// Init will parse the command line flags
	srv.Init()

	// Register handler
	pb.RegisterUserServiceHandler(srv.Server(), &service{repo, tokenService, publisher})

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}

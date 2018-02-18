package main

import (
	"log"
	"os"

	pb "github.com/johynpapin/yonjuuni/user-service/proto/user"
	"github.com/micro/go-micro"
	microclient "github.com/micro/go-micro/client"
	"golang.org/x/net/context"
)

func main() {
	srv := micro.NewService(

		micro.Name("jp.yonjuuni.srv.user-cli"),
		micro.Version("latest"),
	)

	// Init will parse the command line flags.
	srv.Init()

	client := pb.NewUserServiceClient("jp.yonjuuni.srv.user", microclient.DefaultClient)

	username := "johynpapin"
	email := "johynpapin@protonmail.com"
	password := "serpentin22"

	r, err := client.Create(context.TODO(), &pb.User{
		Username: username,
		Email:    email,
		Password: password,
	})
	if err != nil {
		log.Fatalf("Could not create: %v", err)
	}
	log.Printf("Created: %s", r.User.ID)

	getAll, err := client.GetAll(context.Background(), &pb.Request{})
	if err != nil {
		log.Fatalf("Could not list users: %v", err)
	}
	for _, v := range getAll.Users {
		log.Println(v)
	}

	authResponse, err := client.Auth(context.TODO(), &pb.User{
		Email:    email,
		Password: password,
	})

	if err != nil {
		log.Fatalf("Could not authenticate user: %s error: %v\n", email, err)
	}

	log.Printf("Your access token is: %s \n", authResponse.Token)

	os.Exit(0)
}

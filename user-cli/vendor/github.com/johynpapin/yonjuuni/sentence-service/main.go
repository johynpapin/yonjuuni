package main

import (
	pb "github.com/johynpapin/yonjuuni/sentence-service/proto/sentence"
	"github.com/micro/go-micro"
	"fmt"
	"log"
)

func main() {
	// Connect to the database
	db, err := CreateConnection()
	if err != nil {
		log.Panicf("could not connect to database: %v", err)
	}
	defer db.Close()

	db.AutoMigrate(Sentence{})

	repo := &SentenceRepository{db}

	// Create a new service
	srv := micro.NewService(
		micro.Name("jp.yonjuuni.srv.user"),
		micro.Version("latest"),
	)

	// Init will parse the command line flags
	srv.Init()

	// Register handler
	pb.RegisterSentenceServiceHandler(srv.Server(), &service{repo})

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}

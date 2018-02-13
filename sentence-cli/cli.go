package main

import (
	"log"

	pb "github.com/johynpapin/yonjuuni/sentence-service/proto/sentence"
	"golang.org/x/net/context"

	"github.com/micro/go-micro/cmd"
	microclient "github.com/micro/go-micro/client"
)

const (
	address = "localhost:50051"
)

func main() {
	cmd.Init()

	client := pb.NewSentenceServiceClient("jp.yonjuuni.srv.sentence", microclient.DefaultClient)

	insertSentenceRequest := &pb.InsertSentenceRequest{}

	_, err := client.InsertSentence(context.TODO(), insertSentenceRequest)
	if err != nil {
		log.Fatalf("could not insert sentence: %v", err)
	}
	log.Printf("inserted")
}

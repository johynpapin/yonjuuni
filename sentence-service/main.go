package main

import (
	pb "github.com/johynpapin/yonjuuni/sentence-service/proto/sentence"
	"github.com/micro/go-micro"
	"golang.org/x/net/context"
	"fmt"
)

type service struct {
}

func (s *service) InsertSentence(ctx context.Context, req *pb.InsertSentenceRequest, res *pb.InsertSentenceResponse) error {
	res = &pb.InsertSentenceResponse{}

	return nil
}

func main() {
	srv := micro.NewService(
		micro.Name("jp.yonjuuni.srv.sentence"),
		micro.Version("latest"),
	)

	srv.Init()

	pb.RegisterSentenceServiceHandler(srv.Server(), &service{})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}

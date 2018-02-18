package main

import (
	pb "github.com/johynpapin/yonjuuni/sentence-service/proto/sentence"
	"context"
)

type service struct {
	repo Repository
}

func (srv *service) Create(ctx context.Context, req *pb.Sentence, res *pb.Response) error {
	err := srv.repo.Create(req)

	return err
}

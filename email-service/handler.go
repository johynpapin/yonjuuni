package main

import (
	"context"
	pb "github.com/johynpapin/yonjuuni/user-service/proto/user"
	"log"
)

type Service struct {
}

type Subscriber struct{}

func (sub *Subscriber) Process(ctx context.Context, user *pb.User) error {
	err := sendMail(&SmtpTemplateData{"yonjuuni@gmail.com", user.Email, "Olala +1", "Olala +1 !! Vraiment c’est foufou !"})

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

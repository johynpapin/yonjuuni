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
	err := sendMail(Email{"Yonjuuni", "noreply@yonjuuni.jp"}, Email{user.Username, user.Email}, "Bienvenue sur Yonjuuni", "Votre compte a bien été créé.", "Votre compte a bien été <b>créé</b>.")

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

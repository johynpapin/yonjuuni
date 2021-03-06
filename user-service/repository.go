package main

import (
	"github.com/jinzhu/gorm"
	pb "github.com/johynpapin/yonjuuni/user-service/proto/user"
)

type Repository interface {
	GetAll() ([]*pb.User, error)
	Get(user *pb.User) (*pb.User, error)
	Create(user *pb.User) error
}

type UserRepository struct {
	db *gorm.DB
}

func (repo *UserRepository) GetAll() ([]*pb.User, error) {
	var users []*pb.User
	if err := repo.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *UserRepository) Get(user *pb.User) (*pb.User, error) {
	res := &pb.User{}
	if err := repo.db.Where(&pb.User{Username: user.Username, Email: user.Email}).First(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func (repo *UserRepository) Create(user *pb.User) error {
	if err := repo.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

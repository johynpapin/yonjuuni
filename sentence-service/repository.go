package main

import (
	pb "github.com/johynpapin/yonjuuni/sentence-service/proto/sentence"
	"github.com/jinzhu/gorm"
)

type Repository interface {
	Create(sentence *pb.Sentence) error
}

type SentenceRepository struct {
	db *gorm.DB
}

func (repo *SentenceRepository) Create(sentence *pb.Sentence) error {
	err := repo.db.Create(sentence).Error

	return err
}
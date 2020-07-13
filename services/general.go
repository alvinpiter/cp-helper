package services

import "github.com/alvinpiter/cp-helper/entities"

type Service struct {
	CodeforcesRepo entities.Repository
}

func NewService(codeforcesRepo entities.Repository) *Service {
	return &Service{
		CodeforcesRepo: codeforcesRepo,
	}
}

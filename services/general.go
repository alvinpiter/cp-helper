package services

import "github.com/alvinpiter/cp-helper/entities"

type Service struct {
	CodeforcesRepo entities.CodeforcesRepository
}

func NewService(codeforcesRepo entities.CodeforcesRepository) *Service {
	return &Service{
		CodeforcesRepo: codeforcesRepo,
	}
}

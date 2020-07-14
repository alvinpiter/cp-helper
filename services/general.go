package services

import (
	"github.com/alvinpiter/cp-helper/entities"
	"github.com/alvinpiter/cp-helper/repos/nethttp"
)

type Service struct {
	CodeforcesRepo entities.Repository
	AtCoderRepo    entities.Repository
}

func NewService() *Service {
	cfRepo := nethttp.NewCodeforcesRepository()
	atcRepo := nethttp.NewAtCoderRespository()

	return &Service{
		CodeforcesRepo: cfRepo,
		AtCoderRepo:    atcRepo,
	}
}

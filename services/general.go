package services

import (
	"github.com/alvinpiter/cp-helper/entities"
	"github.com/alvinpiter/cp-helper/repos/nethttp"
)

type Service struct {
	CodeforcesRepo entities.Repository
}

func NewService() *Service {
	cfRepo := nethttp.NewCodeforcesRepository()
	return &Service{
		CodeforcesRepo: cfRepo,
	}
}

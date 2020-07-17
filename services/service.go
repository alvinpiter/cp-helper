package services

import (
	"github.com/alvinpiter/cp-helper/entities"
	"github.com/alvinpiter/cp-helper/repos/atcoder"
	"github.com/alvinpiter/cp-helper/repos/codeforces"
)

type Service struct {
	CodeforcesRepo entities.Repository
	AtCoderRepo    entities.Repository
}

func New() *Service {
	cfRepo := codeforces.NewRepository()
	atcRepo := atcoder.NewRepository()

	return &Service{
		CodeforcesRepo: cfRepo,
		AtCoderRepo:    atcRepo,
	}
}

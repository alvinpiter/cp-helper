package api

import "github.com/alvinpiter/cp-helper/entities"

type CompareRequest struct {
	OnlineJudge *string                   `json:"online_judge"`
	Handle      *string                   `json:"handle"`
	RivalHandle *string                   `json:"rival_handle"`
	Filter      *entities.FilterParameter `json:"filter"`
}

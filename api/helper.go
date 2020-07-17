package api

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	ErrEmptyOnlineJudge      = errors.New("online_judge is required")
	ErrEmptyHandle           = errors.New("handle is required")
	ErrEmptyRivalHandle      = errors.New("rival_handle is required")
	ErrInvalidTagsFilterMode = errors.New("tags mode is either `and` or `or`")
	ErrEmptyTagsFilterValues = errors.New("tags values is required")
)

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	jsonData, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(jsonData)
}

func NormalizeCompareRequest(r *http.Request) (*CompareRequest, error) {
	cp := &CompareRequest{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(cp)
	if err != nil {
		return nil, err
	}

	if cp.OnlineJudge == nil {
		return nil, ErrEmptyOnlineJudge
	}

	if cp.Handle == nil {
		return nil, ErrEmptyHandle
	}

	if cp.RivalHandle == nil {
		return nil, ErrEmptyRivalHandle
	}

	if cp.Filter != nil {
		if *cp.OnlineJudge == "atcoder" {
			cp.Filter.Tags = nil
		} else {
			tagsFilter := cp.Filter.Tags

			if tagsFilter.Mode != "and" && tagsFilter.Mode != "or" {
				return nil, ErrInvalidTagsFilterMode
			}

			if tagsFilter.Values == nil {
				return nil, ErrEmptyTagsFilterValues
			}
		}
	}

	return cp, nil
}

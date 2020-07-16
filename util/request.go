package util

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/alvinpiter/cp-helper/entities"
)

/*
NormalizeRequestBody normalizes request body and performs some validation on i.
*/
func NormalizeRequestBody(oj string, body io.Reader) (*entities.RequestParameter, error) {
	rp := &entities.RequestParameter{}

	switch oj {
	case "codeforces":
		rp.Filter = &entities.CodeforcesFilterParameter{}
	case "atcoder":
		rp.Filter = &entities.AtCoderFilterParameter{}
	}

	decoder := json.NewDecoder(body)
	err := decoder.Decode(rp)

	if err != nil {
		return nil, err
	}

	//Validate parameters

	if rp.Handle == "" {
		return nil, errors.New("Handle can't be empty")
	}

	if rp.RivalHandle == "" {
		return nil, errors.New("Rival handle can't be empty")
	}

	tagsFilter := rp.Filter.GetTagsFilterParameter()

	if tagsFilter != nil {
		if tagsFilter.Mode != "and" && tagsFilter.Mode != "or" {
			return nil, errors.New("Tags filter mode is either `and` or `or`")
		}

		if tagsFilter.Values == nil {
			return nil, errors.New("Tags can't be nil")
		}
	}

	return rp, nil
}

package handlers

import (
	"errors"
	"net/http"

	"github.com/alvinpiter/cp-helper/entities"
	"github.com/alvinpiter/cp-helper/util"
)

func CompareAtCoder(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		jsonError(w, errors.New("Invalid HTTP method"))
		return
	}

	params, err := util.NormalizeRequestBody("atcoder", req.Body)
	if err != nil {
		jsonError(w, err)
		return
	}

	handle := params.Handle
	rivalHandle := params.RivalHandle
	diffs, err := service.Compare("atcoder", handle, rivalHandle)
	if err != nil {
		jsonError(w, err)
		return
	}

	diffs = service.ApplyProblemFilter(diffs, params.Filter)

	jsonResponse(w, http.StatusOK, entities.NewCompareResponse(diffs))
}

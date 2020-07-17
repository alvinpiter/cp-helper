package api_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alvinpiter/cp-helper/api"
	"github.com/stretchr/testify/assert"
)

func TestHealthz(t *testing.T) {
	req, err := http.NewRequest("GET", "/healthz", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(api.HealthzHandler)

	handler.ServeHTTP(recorder, req)

	assert.Equal(t, 200, recorder.Code)
}

func TestCodeforcesProblemTags(t *testing.T) {
	req, err := http.NewRequest("GET", "/codeforces-problem-tags", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(api.CodeforcesProblemTagsHandler)

	handler.ServeHTTP(recorder, req)

	respObj := api.CodeforcesProblemTagsResponse{}
	decoder := json.NewDecoder(recorder.Body)
	err = decoder.Decode(&respObj)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 200, recorder.Code)
	assert.Equal(t, 37, len(respObj.ProblemTags))
}

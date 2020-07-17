package api_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/alvinpiter/cp-helper/api"
	"github.com/alvinpiter/cp-helper/entities"
	"github.com/alvinpiter/cp-helper/entities/mocks"
	"github.com/stretchr/testify/assert"
)

func TestHealthz(t *testing.T) {
	app := api.New()

	req, err := http.NewRequest("GET", "/healthz", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(app.HealthzHandler)

	handler.ServeHTTP(recorder, req)

	assert.Equal(t, 200, recorder.Code)
}

func TestCodeforcesProblemTags(t *testing.T) {
	app := api.New()

	req, err := http.NewRequest("GET", "/codeforces-problem-tags", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(app.CodeforcesProblemTagsHandler)

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

func TestCompare(t *testing.T) {
	problem := entities.Problem{
		ID:     "A",
		Name:   "Problem A",
		Rating: 2000,
		Tags:   []string{"implementation"},
		URL:    "https://codeforces.com",
	}

	svcMock := new(mocks.Service)
	svcMock.On(
		"CompareWithFilter",
		"codeforces",
		"handle",
		"rival_handle",
		mock.Anything).Return([]entities.Problem{problem}, nil)

	app := api.New()
	app.Service = svcMock

	bodyJSON := `{
		"online_judge": "codeforces",
		"handle": "handle",
		"rival_handle": "rival_handle"
	}`

	req, err := http.NewRequest("POST", "/compare", strings.NewReader(bodyJSON))
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(app.CompareHandler)

	handler.ServeHTTP(recorder, req)

	respObj := api.CompareResponse{}
	decoder := json.NewDecoder(recorder.Body)
	err = decoder.Decode(&respObj)
	if err != nil {
		t.Fatal(err)
	}

	p := respObj.Problems[0]

	assert.Equal(t, 1, len(respObj.Problems))
	assert.Equal(t, "A", p.ID)
	assert.Equal(t, "Problem A", p.Name)
	assert.Equal(t, 2000, p.Rating)
	assert.Equal(t, 1, len(p.Tags))
	assert.Equal(t, "implementation", p.Tags[0])
	assert.Equal(t, "https://codeforces.com", p.URL)
}

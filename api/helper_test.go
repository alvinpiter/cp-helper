package api_test

import (
	"net/http"
	"strings"
	"testing"

	"github.com/alvinpiter/cp-helper/api"
	"github.com/stretchr/testify/assert"
)

func TestNormalizeCompareRequest(t *testing.T) {
	body1 := `{
		"handle": "handle",
		"rival_handle": "rival_handle"
	}`

	body2 := `{
		"online_judge": "codeforces",
		"rival_handle": "rival_handle"
	}`

	body3 := `{
		"online_judge": "codeforces",
		"handle": "handle"
	}`

	body4 := `{
		"online_judge": "codeforces",
		"handle": "handle",
		"rival_handle": "rival_handle",
		"filter": {
			"rating": {
				"minimum": 1000,
				"maximum": 2000
			},
			"tags": {
				"mode": "and",
				"values": ["implementation"]
			}
		}
	}`

	body5 := `{
		"online_judge": "codeforces",
		"handle": "handle",
		"rival_handle": "rival_handle",
		"filter": {
			"tags": {
				"values": ["implementation"]
			}
		}
	}`

	body6 := `{
		"online_judge": "codeforces",
		"handle": "handle",
		"rival_handle": "rival_handle",
		"filter": {
			"tags": {
				"mode": "and"
			}
		}
	}`

	body7 := `{
		"online_judge": "atcoder",
		"handle": "handle",
		"rival_handle": "rival_handle",
		"filter": {
			"tags": {
				"mode": "and",
				"values": ["implementation", "math"]
			}
		}
	}`

	req1, _ := http.NewRequest("POST", "/compare", strings.NewReader(body1))
	req2, _ := http.NewRequest("POST", "/compare", strings.NewReader(body2))
	req3, _ := http.NewRequest("POST", "/compare", strings.NewReader(body3))
	req4, _ := http.NewRequest("POST", "/compare", strings.NewReader(body4))
	req5, _ := http.NewRequest("POST", "/compare", strings.NewReader(body5))
	req6, _ := http.NewRequest("POST", "/compare", strings.NewReader(body6))
	req7, _ := http.NewRequest("POST", "/compare", strings.NewReader(body7))

	_, err1 := api.NormalizeCompareRequest(req1)
	assert.Equal(t, api.ErrEmptyOnlineJudge, err1)

	_, err2 := api.NormalizeCompareRequest(req2)
	assert.Equal(t, api.ErrEmptyHandle, err2)

	_, err3 := api.NormalizeCompareRequest(req3)
	assert.Equal(t, api.ErrEmptyRivalHandle, err3)

	res4, err4 := api.NormalizeCompareRequest(req4)
	assert.Nil(t, err4)
	assert.Equal(t, "codeforces", *res4.OnlineJudge)
	assert.Equal(t, "handle", *res4.Handle)
	assert.Equal(t, "rival_handle", *res4.RivalHandle)
	assert.Equal(t, 1000, *res4.Filter.Rating.Minimum)
	assert.Equal(t, 2000, *res4.Filter.Rating.Maximum)
	assert.Equal(t, "and", res4.Filter.Tags.Mode)
	assert.Equal(t, 1, len(res4.Filter.Tags.Values))
	assert.Equal(t, "implementation", res4.Filter.Tags.Values[0])

	_, err5 := api.NormalizeCompareRequest(req5)
	assert.Equal(t, api.ErrInvalidTagsFilterMode, err5)

	_, err6 := api.NormalizeCompareRequest(req6)
	assert.Equal(t, api.ErrEmptyTagsFilterValues, err6)

	res7, _ := api.NormalizeCompareRequest(req7)
	//TODO: Can we use assert instead of this?
	if res7.Filter.Tags != nil {
		t.Error()
	}
}

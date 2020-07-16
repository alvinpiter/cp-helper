package util_test

import (
	"strings"
	"testing"

	"github.com/alvinpiter/cp-helper/util"
	"github.com/stretchr/testify/assert"
)

func TestNormalizeRequestBody(t *testing.T) {
	body1Json := `
		{
			"handle": "alvinpiter",
			"rival_handle": "chokudai",
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
		}
	`

	body2Json := `
		{
			"handle": "alvinpiter"
		}
	`

	body3Json := `
		{
			"rival_handle": "chokudai"
		}
	`

	body4Json := `
		{
			"handle": "alvinpiter",
			"rival_handle": "chokudai",
			"filter": {
				"tags": {
					"mode": "lol"
				}
			}
		}
	`

	body5Json := `
		{
			"handle": "alvinpiter",
			"rival_handle": "chokudai",
			"filter": {
				"tags": {
					"mode": "and"
				}
			}
		}
	`

	body6Json := `
		{
			"handle": "alvinpiter",
			"rival_handle": "chokudai"
		}
	`
	body1 := strings.NewReader(body1Json)
	body2 := strings.NewReader(body2Json)
	body3 := strings.NewReader(body3Json)
	body4 := strings.NewReader(body4Json)
	body5 := strings.NewReader(body5Json)
	body6 := strings.NewReader(body6Json)

	res1, err1 := util.NormalizeRequestBody("codeforces", body1)
	assert.Nil(t, err1)
	assert.Equal(t, "alvinpiter", res1.Handle)
	assert.Equal(t, "chokudai", res1.RivalHandle)
	assert.Equal(t, 1000, res1.Filter.GetRatingFilterParameter().Minimum)
	assert.Equal(t, 2000, res1.Filter.GetRatingFilterParameter().Maximum)
	assert.Equal(t, "and", res1.Filter.GetTagsFilterParameter().Mode)
	assert.Equal(t, 1, len(res1.Filter.GetTagsFilterParameter().Values))
	assert.Equal(t, "implementation", res1.Filter.GetTagsFilterParameter().Values[0])

	_, err2 := util.NormalizeRequestBody("codeforces", body2)
	assert.Equal(t, "Rival handle can't be empty", err2.Error())

	_, err3 := util.NormalizeRequestBody("codeforces", body3)
	assert.Equal(t, "Handle can't be empty", err3.Error())

	_, err4 := util.NormalizeRequestBody("codeforces", body4)
	assert.Equal(t, "Tags filter mode is either `and` or `or`", err4.Error())

	_, err5 := util.NormalizeRequestBody("codeforces", body5)
	assert.Equal(t, "Tags can't be nil", err5.Error())

	_, err6 := util.NormalizeRequestBody("atcoder", body6)
	assert.Nil(t, err6)
	//TODO: Assert tag filter is nil
}

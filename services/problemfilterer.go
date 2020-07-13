/*
TODO:
* DRY mapToStruct
* DRY filterFuncBuilder
* Check if filterFuncBuilder is called too many times
*/

package services

import (
	"github.com/alvinpiter/cp-helper/entities"
	"github.com/mitchellh/mapstructure"
)

type FilterFunc func(entities.Problem) bool

type tagFilterParameter struct {
	Mode   string
	Values []string
}

type ratingFilterParameter struct {
	Minimum int
	Maximum int
}

type idFilterParameter struct {
	Mode   string
	Values []string
}

type filterParameter struct {
	TagFilter    *tagFilterParameter
	RatingFilter *ratingFilterParameter
	IdFilter     *idFilterParameter
}

func (s *Service) ApplyProblemFilter(problems []entities.Problem, params map[string]interface{}) []entities.Problem {
	filterParam := mapToStruct(params)
	filterFunc := filterFuncBuilder(filterParam)

	return doApplyProblemFilter(problems, filterFunc)
}

func doApplyProblemFilter(problems []entities.Problem, filterFunc FilterFunc) []entities.Problem {
	result := []entities.Problem{}
	for _, problem := range problems {
		if filterFunc(problem) {
			result = append(result, problem)
		}
	}

	return result
}

func mapToStruct(params map[string]interface{}) *filterParameter {
	fp := &filterParameter{}

	if params["tag"] != nil {
		mp, ok := params["tag"].(map[string]interface{})
		if ok {
			tagFilter := &tagFilterParameter{}
			err := mapstructure.Decode(mp, tagFilter)

			if err == nil && (tagFilter.Mode == "and" || tagFilter.Mode == "or") {
				fp.TagFilter = tagFilter
			}
		}
	}

	if params["rating"] != nil {
		mp, ok := params["rating"].(map[string]int)
		if ok {
			ratingFilter := &ratingFilterParameter{}
			err := mapstructure.Decode(mp, ratingFilter)

			if err == nil {
				fp.RatingFilter = ratingFilter
			}
		}
	}

	if params["id"] != nil {
		mp, ok := params["id"].(map[string]interface{})
		if ok {
			idFilter := &idFilterParameter{}
			err := mapstructure.Decode(mp, idFilter)

			if err == nil && (idFilter.Mode == "exclude" || idFilter.Mode == "include") {
				fp.IdFilter = idFilter
			}
		}
	}

	return fp
}

func filterFuncBuilder(fp *filterParameter) FilterFunc {
	funcs := []FilterFunc{}

	if fp.TagFilter != nil {
		funcs = append(funcs, tagFilterFuncBuilder(fp.TagFilter))
	}

	if fp.RatingFilter != nil {
		funcs = append(funcs, ratingFilterFuncBuilder(fp.RatingFilter))
	}

	if fp.IdFilter != nil {
		funcs = append(funcs, idFilterFuncBuilder(fp.IdFilter))
	}

	return func(p entities.Problem) bool {
		valid := true
		for _, f := range funcs {
			valid = (valid && f(p))
		}

		return valid
	}
}

func tagFilterFuncBuilder(tagFilter *tagFilterParameter) FilterFunc {
	if tagFilter.Mode == "or" {
		return func(p entities.Problem) bool {
			tagMap := make(map[string]bool)
			for _, tag := range tagFilter.Values {
				tagMap[tag] = true
			}

			for _, tag := range p.GetTags() {
				if _, exist := tagMap[tag]; exist {
					return true
				}
			}

			return true
		}
	}

	//"and" mode
	return func(p entities.Problem) bool {
		tagMap := make(map[string]bool)
		for _, tag := range p.GetTags() {
			tagMap[tag] = true
		}

		for _, tag := range tagFilter.Values {
			if _, exist := tagMap[tag]; !exist {
				return false
			}
		}

		return true
	}
}

func ratingFilterFuncBuilder(ratingFilter *ratingFilterParameter) FilterFunc {
	minRating := ratingFilter.Minimum
	maxRating := ratingFilter.Maximum

	return func(p entities.Problem) bool {
		return p.GetRating() >= minRating && p.GetRating() <= maxRating
	}
}

func idFilterFuncBuilder(idFilter *idFilterParameter) FilterFunc {
	idMap := make(map[string]bool)
	for _, id := range idFilter.Values {
		idMap[id] = true
	}

	if idFilter.Mode == "exclude" {
		return func(p entities.Problem) bool {
			if _, exist := idMap[p.GetID()]; !exist {
				return true
			}

			return false
		}
	}

	//Implement "include" mode if necessary
	return func(entities.Problem) bool {
		return true
	}
}

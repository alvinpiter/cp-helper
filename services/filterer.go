/*
TODO:
* DRY mapToStruct
* DRY filterFuncBuilder
* Check if filterFuncBuilder is called too many times
*/

package services

import (
	"github.com/alvinpiter/cp-helper/entities"
)

type FilterFunc func(entities.Problem) bool

func ApplyProblemFilter(problems []entities.Problem, fp *entities.FilterParameter) []entities.Problem {
	filterFunc := filterFuncBuilder(fp)
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

func filterFuncBuilder(fp *entities.FilterParameter) FilterFunc {
	funcs := []FilterFunc{}

	if fp.Rating != nil {
		funcs = append(funcs, ratingFilterFuncBuilder(fp.Rating))
	}

	if fp.Tags != nil {
		funcs = append(funcs, tagsFilterFuncBuilder(fp.Tags))
	}

	return func(p entities.Problem) bool {
		valid := true
		for _, f := range funcs {
			valid = (valid && f(p))
		}

		return valid
	}
}

func tagsFilterFuncBuilder(tagsFilter *entities.TagsFilterParameter) FilterFunc {
	if tagsFilter.Mode == "or" {
		return func(p entities.Problem) bool {
			tagMap := make(map[string]bool)
			for _, tag := range tagsFilter.Values {
				tagMap[tag] = true
			}

			for _, tag := range p.Tags {
				if _, exist := tagMap[tag]; exist {
					return true
				}
			}

			return false
		}
	}

	//"and" mode
	return func(p entities.Problem) bool {
		tagMap := make(map[string]bool)
		for _, tag := range p.Tags {
			tagMap[tag] = true
		}

		for _, tag := range tagsFilter.Values {
			if _, exist := tagMap[tag]; !exist {
				return false
			}
		}

		return true
	}
}

func ratingFilterFuncBuilder(ratingFilter *entities.RatingFilterParameter) FilterFunc {
	var minRating, maxRating int

	if ratingFilter.Minimum == nil {
		minRating = 0
	} else {
		minRating = *ratingFilter.Minimum
	}

	if ratingFilter.Maximum == nil {
		maxRating = 5000
	} else {
		maxRating = *ratingFilter.Maximum
	}

	return func(p entities.Problem) bool {
		return p.Rating >= minRating && p.Rating <= maxRating
	}
}

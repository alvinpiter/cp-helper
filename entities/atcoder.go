package entities

import "fmt"

var problemURLFormat = "https://atcoder.jp/contests/%s/tasks/%s"

type AtCoderProblem struct {
	ID         string `json:"id"`
	ContestID  string `json:"contest_id"`
	Title      string `json:"title"`
	Difficulty float64
}

func (p *AtCoderProblem) GetID() string {
	return p.ID
}

func (p *AtCoderProblem) GetName() string {
	return p.Title
}

func (p *AtCoderProblem) GetRating() int {
	rating := int(p.Difficulty)
	if rating < 0 {
		return 0
	}

	return rating
}

func (p *AtCoderProblem) GetTags() []string {
	return []string{}
}

func (p *AtCoderProblem) GetURL() string {
	return fmt.Sprintf(problemURLFormat, p.ContestID, p.ID)
}

type AtCoderProblemDifficulty struct {
	Difficulty float64 `json:"difficulty"`
}

type AtCoderSubmission struct {
	ProblemID string `json:"problem_id"`
	Problem   *AtCoderProblem
	Result    string `json:"result"`
}

func (s *AtCoderSubmission) GetProblem() Problem {
	return s.Problem
}

func (s *AtCoderSubmission) IsAccepted() bool {
	if s.Result == "AC" {
		return true
	}

	return false
}

/*
AtCoderFilterParameter is a struct that represents problem filter parameters for AtCoder
*/
type AtCoderFilterParameter struct {
	RatingFilter *RatingFilterParameter `json:"rating"`
}

func (a *AtCoderFilterParameter) GetTagsFilterParameter() *TagsFilterParameter {
	return &TagsFilterParameter{
		Mode:   "and",
		Values: []string{},
	}
}

func (a *AtCoderFilterParameter) GetRatingFilterParameter() *RatingFilterParameter {
	return a.RatingFilter
}

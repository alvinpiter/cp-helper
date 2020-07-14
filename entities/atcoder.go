package entities

import "fmt"

var problemURLFormat = "https://atcoder.jp/contests/%s/tasks/%s"

type AtCoderProblem struct {
	ID         string  `json:"id"`
	ContestID  string  `json:"contest_id"`
	Title      string  `json:"title"`
	Difficulty float64 `json:"difficulty"`
}

func (p *AtCoderProblem) GetID() string {
	return p.ID
}

func (p *AtCoderProblem) GetName() string {
	return p.Title
}

func (p *AtCoderProblem) GetRating() int {
	return int(p.Difficulty)
}

func (p *AtCoderProblem) GetTags() []string {
	return []string{}
}

func (p *AtCoderProblem) GetURL() string {
	return fmt.Sprintf(problemURLFormat, p.ContestID, p.ID)
}

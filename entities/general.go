package entities

/*
Problem is an interface that represents an online judge problem
*/
type Problem interface {
	GetID() string
	GetName() string
	GetRating() int
	GetTags() []string
	GetURL() string
}

/*
Submission is an interface that represents an online judge submission
*/
type Submission interface {
	GetProblem() Problem
	IsAccepted() bool
}

/*
Repository is an interface that represents an object that
communicates with online judge API
*/
type Repository interface {
	GetProblems() ([]Problem, error)
	GetSubmissions(string) ([]Submission, error)
}

type RatingFilterParameter struct {
	Minimum int `json:"minimum"`
	Maximum int `json:"maximum"`
}

type TagsFilterParameter struct {
	Mode   string   `json:"mode"`
	Values []string `json:"values"`
}

type FilterParameter interface {
	GetRatingFilterParameter() *RatingFilterParameter
	GetTagsFilterParameter() *TagsFilterParameter
}

type RequestParameter struct {
	Handle      string          `json:"handle"`
	RivalHandle string          `json:"rival_handle"`
	Filter      FilterParameter `json:"filter"`
}

package entities

type Problem struct {
	ID     string   `json:"id"`
	Name   string   `json:"name"`
	Rating int      `json:"rating"`
	Tags   []string `json:"tags"`
	URL    string   `json:"url"`
}

type Submission struct {
	Problem    Problem
	IsAccepted bool
}

type Service interface {
	CompareWithFilter(string, string, string, *FilterParameter) ([]Problem, error)
}

/*
Repository is an interface that represents an object that
communicates with online judge API
*/
type Repository interface {
	GetSubmissions(string) ([]Submission, error)
}

type RatingFilterParameter struct {
	Minimum *int `json:"minimum"`
	Maximum *int `json:"maximum"`
}

type TagsFilterParameter struct {
	Mode   string   `json:"mode"`
	Values []string `json:"values"`
}

type FilterParameter struct {
	Rating *RatingFilterParameter `json:"rating"`
	Tags   *TagsFilterParameter   `json:"tags"`
}

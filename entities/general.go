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
	Minimum *int `json:"minimum"`
	Maximum *int `json:"maximum"`
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

type ProblemResponse struct {
	ID     string   `json:"id"`
	Name   string   `json:"name"`
	Rating int      `json:"rating"`
	Tags   []string `json:"tags"`
	URL    string   `json:"url"`
}

type CompareResponse struct {
	Problems []ProblemResponse `json:"problems"`
}

func NewCompareResponse(problems []Problem) CompareResponse {
	res := []ProblemResponse{}
	for _, p := range problems {
		pr := ProblemResponse{
			ID:     p.GetID(),
			Name:   p.GetName(),
			Rating: p.GetRating(),
			Tags:   p.GetTags(),
			URL:    p.GetURL(),
		}

		res = append(res, pr)
	}

	return CompareResponse{
		Problems: res,
	}
}

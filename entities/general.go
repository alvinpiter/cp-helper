package entities

type Problem struct {
	ID     string
	Name   string
	Rating int
	Tags   []string
	URL    string
}

type Submission struct {
	Problem    Problem
	IsAccepted bool
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
	// for _, p := range problems {
	// 	pr := ProblemResponse{
	// 		ID:     p.GetID(),
	// 		Name:   p.GetName(),
	// 		Rating: p.GetRating(),
	// 		Tags:   p.GetTags(),
	// 		URL:    p.GetURL(),
	// 	}

	// 	res = append(res, pr)
	// }

	return CompareResponse{
		Problems: res,
	}
}

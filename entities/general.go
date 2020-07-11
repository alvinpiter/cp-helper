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
Repository is an interface that represents an object that communicates with
online judge API.
*/
type Repository interface {
	GetSubmissions(string) ([]Submission, error)
	GetProblems() ([]Problem, error)
}

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

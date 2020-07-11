package entities

/*
Problem is an interface that represents an online judge problem
*/
type Problem interface {
	getID() string
	getName() string
	getRating() int
	getTags() []string
	getURL() string
}

/*
Submission is an interface that represents an online judge submission
*/
type Submission interface {
	GetProblemID() string
	IsAccepted() bool
}

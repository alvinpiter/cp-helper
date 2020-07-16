package entities

import (
	"fmt"
)

var (
	gymProblemURLFormat        = "https://codeforces.com/gym/%d/problem/%s"
	problemsetProblemURLFormat = "https://codeforces.com/contest/%d/problem/%s"
)

/*
CodeforcesProblem is a struct that represents a Problem object in
Codeforces (https://codeforces.com/apiHelp/objects#Problem).
Some fields are not included because they are not needed.
*/
type CodeforcesProblem struct {
	ContestID int      `json:"contestId"`
	Index     string   `json:"index"`
	Name      string   `json:"name"`
	Rating    int      `json:"rating"`
	Tags      []string `json:"tags"`
}

func (c *CodeforcesProblem) GetID() string {
	return fmt.Sprintf("%d%s", c.ContestID, c.Index)
}

func (c *CodeforcesProblem) GetName() string {
	return c.Name
}

func (c *CodeforcesProblem) GetRating() int {
	return c.Rating
}

func (c *CodeforcesProblem) GetTags() []string {
	return c.Tags
}

func (c *CodeforcesProblem) GetURL() string {
	if c.ContestID >= 100000 {
		return fmt.Sprintf(gymProblemURLFormat, c.ContestID, c.Index)
	}

	return fmt.Sprintf(problemsetProblemURLFormat, c.ContestID, c.Index)
}

/*
CodeforcesSubmission is a struct that represents a Submission object
in Codeforces (https://codeforces.com/apiHelp/objects#Submission).
Some fields are not included because they are not needed.
*/
type CodeforcesSubmission struct {
	Problem *CodeforcesProblem `json:"problem"`
	Verdict string             `json:"verdict"`
}

func (c *CodeforcesSubmission) GetProblem() Problem {
	return c.Problem
}

func (c *CodeforcesSubmission) IsAccepted() bool {
	if c.Verdict == "OK" {
		return true
	}

	return false
}

/*
CodeforcesFilterParameter is a struct that represents problem filter parameters for Codeforces
*/
type CodeforcesFilterParameter struct {
	TagsFilter   *TagsFilterParameter   `json:"tags"`
	RatingFilter *RatingFilterParameter `json:"rating"`
}

func (c *CodeforcesFilterParameter) GetTagsFilterParameter() *TagsFilterParameter {
	return c.TagsFilter
}

func (c *CodeforcesFilterParameter) GetRatingFilterParameter() *RatingFilterParameter {
	return c.RatingFilter
}

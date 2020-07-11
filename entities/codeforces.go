package entities

import "fmt"

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
	ContestID int
	Index     string
	Name      string
	Rating    int
	Tags      []string
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
	Problem CodeforcesProblem
	Verdict string
}

func (c *CodeforcesSubmission) GetProblemID() string {
	return c.Problem.GetID()
}

func (c *CodeforcesSubmission) IsAccepted() bool {
	if c.Verdict == "OK" {
		return true
	}

	return false
}

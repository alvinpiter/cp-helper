package services

import (
	"errors"

	"github.com/alvinpiter/cp-helper/entities"
)

func (s *Service) GetAcceptedProblems(oj, handle string) ([]entities.Problem, error) {
	var submissions []entities.Submission
	var err error

	switch oj {
	case "codeforces":
		submissions, err = s.CodeforcesRepo.GetSubmissions(handle)
	default:
		return nil, errors.New("Unknown online judge")
	}

	if err != nil {
		return nil, err
	}

	acProblems := []entities.Problem{}
	seenID := make(map[string]bool) //To avoid duplicates
	for _, submission := range submissions {
		if submission.IsAccepted() {
			problemID := submission.GetProblem().GetID()
			if _, seen := seenID[problemID]; !seen {
				acProblems = append(acProblems, submission.GetProblem())
				seenID[problemID] = true
			}
		}
	}

	return acProblems, nil
}

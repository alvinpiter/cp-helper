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
	case "atcoder":
		submissions, err = s.AtCoderRepo.GetSubmissions(handle)
	default:
		return nil, errors.New("Unknown online judge")
	}

	if err != nil {
		return nil, err
	}

	acProblems := []entities.Problem{}
	seenID := make(map[string]bool) //To avoid duplicates
	for _, submission := range submissions {
		if submission.IsAccepted {
			problemID := submission.Problem.ID
			if _, seen := seenID[problemID]; !seen {
				acProblems = append(acProblems, submission.Problem)
				seenID[problemID] = true
			}
		}
	}

	return acProblems, nil
}

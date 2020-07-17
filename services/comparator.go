package services

import (
	"errors"

	"github.com/alvinpiter/cp-helper/entities"
)

type channelItem struct {
	Data  []entities.Problem
	Error error
}

/*
Compare returns a list of problems that is solved by handle2 but
not solved by handle1
*/
func (s *Service) Compare(oj, handle1, handle2 string) ([]entities.Problem, error) {
	ch1 := make(chan channelItem)
	ch2 := make(chan channelItem)

	go s.getAcceptedProblemsConcurrently(oj, handle1, ch1)
	go s.getAcceptedProblemsConcurrently(oj, handle2, ch2)

	channelItem1 := <-ch1
	acProblems1, err := channelItem1.Data, channelItem1.Error
	if err != nil {
		return nil, err
	}

	channelItem2 := <-ch2
	acProblems2, err := channelItem2.Data, channelItem2.Error
	if err != nil {
		return nil, err
	}

	acProblemIdsMap1 := make(map[string]bool)
	for _, p := range acProblems1 {
		acProblemIdsMap1[p.ID] = true
	}

	diffs := []entities.Problem{}
	for _, p := range acProblems2 {
		if _, exist := acProblemIdsMap1[p.ID]; exist == false {
			diffs = append(diffs, p)
		}
	}

	return diffs, nil
}

func (s *Service) getAcceptedProblemsConcurrently(oj, handle string, ch chan channelItem) {
	problems, err := s.getAcceptedProblems(oj, handle)
	ch <- channelItem{
		problems,
		err,
	}
}

func (s *Service) getAcceptedProblems(oj, handle string) ([]entities.Problem, error) {
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

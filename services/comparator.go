package services

import "github.com/alvinpiter/cp-helper/entities"

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

	acProblemIds1 := []string{}
	for _, p := range acProblems1 {
		acProblemIds1 = append(acProblemIds1, p.GetID())
	}

	filter := map[string]interface{}{
		"id": map[string]interface{}{
			"mode":   "exclude",
			"values": acProblemIds1,
		},
	}

	return s.ApplyProblemFilter(acProblems2, filter), nil
}

func (s *Service) getAcceptedProblemsConcurrently(oj, handle string, ch chan channelItem) {
	problems, err := s.GetAcceptedProblems(oj, handle)
	ch <- channelItem{
		problems,
		err,
	}
}

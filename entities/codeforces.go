package entities

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

type CodeforcesProblemTagsResponse struct {
	ProblemTags []string `json:"problem_tags"`
}

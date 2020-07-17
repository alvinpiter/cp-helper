package entities

/*
AtCoderFilterParameter is a struct that represents problem filter parameters for AtCoder
*/
type AtCoderFilterParameter struct {
	RatingFilter *RatingFilterParameter `json:"rating"`
}

func (a *AtCoderFilterParameter) GetTagsFilterParameter() *TagsFilterParameter {
	return nil
}

func (a *AtCoderFilterParameter) GetRatingFilterParameter() *RatingFilterParameter {
	return a.RatingFilter
}

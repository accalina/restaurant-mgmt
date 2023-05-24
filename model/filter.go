package model

type Filter struct {
	Limit     int    `json:"limit" default:"10"`
	Page      int    `json:"page" default:"1"`
	Offset    int    `json:"offset"`
	Search    string `json:"search,omitempty"`
	OrderBy   string `json:"orderBy,omitempty" default:"updated_at"`
	Sort      string `json:"sort,omitempty" default:"desc" lower:"true"`
	StartDate string `json:"startDate,omitempty" time_format:"2006-01-02"`
	EndDate   string `json:"endDate,omitempty" time_format:"2006-01-02"`
}

func (f *Filter) CalculateOffset() int {
	f.Offset = (f.Page - 1) * f.Limit
	return f.Offset
}

func (f *Filter) SetDefaultOrderBy() {
	if f.OrderBy == "" {
		f.OrderBy = "updated_at"
	}
}

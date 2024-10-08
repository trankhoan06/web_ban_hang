package model

type Paging struct {
	Page  int   `json:"page" form:"page"`
	Limit int   `json:"limit" form:"limit"`
	Total int64 `json:"total" form:"-"`
}

func (p *Paging) Process() {
	if p.Limit == 0 || p.Limit > 10 {
		p.Limit = 10
	}
	if p.Page < 1 {
		p.Page = 1
	}
}

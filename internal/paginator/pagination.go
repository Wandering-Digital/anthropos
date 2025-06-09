package paginator

import (
	"math"
	"net/url"
)

type Pagination struct {
	CurrentPage int `json:"current_page"`
	PerPage     int `json:"per_page"`
	TotalPage   int `json:"total_page"`
	Total       int `json:"total"`
}

func NewPagination(url *url.URL, perPage int, total int) *Pagination {
	if perPage < 1 {
		perPage = 1
	}

	if total < 0 {
		total = 1
	}

	return &Pagination{
		CurrentPage: 1,
		PerPage:     perPage,
		TotalPage:   int(math.Ceil(float64(total) / float64(perPage))),
		Total:       total,
	}
}

func (p *Pagination) Offset() int {
	return (p.CurrentPage - 1) * p.PerPage
}

func (p *Pagination) Limit() int {
	return p.PerPage
}

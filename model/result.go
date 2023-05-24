package model

import "math"

type Meta struct {
	Page         int `json:"page"`
	Limit        int `json:"limit"`
	TotalRecords int `json:"totalRecords"`
	TotalPages   int `json:"totalPages"`
}

func NewMeta(page, limit, totalRecords int) (m Meta) {
	m.Page, m.Limit, m.TotalRecords = page, limit, totalRecords
	m.CalculatePages()
	return
}

func (m *Meta) CalculatePages() {
	m.TotalPages = int(math.Ceil(float64(m.TotalRecords) / float64(m.Limit)))
}

package response

import "yasser-backend/pkg/pagination"

type PaginationMeta struct {
	CurrentPage int   `json:"currentPage"`
	PerPage     int   `json:"perPage"`
	Total       int64 `json:"total"`
	LastPage    int   `json:"lastPage"`
	HasNext     bool  `json:"hasNext"`
	HasPrev     bool  `json:"hasPrev"`
}

type PaginatedResponse struct {
	Items      interface{}     `json:"items"`
	Pagination *PaginationMeta `json:"pagination"`
}

func NewPaginatedResponse(items interface{}, meta *PaginationMeta) *PaginatedResponse {
	return &PaginatedResponse{
		Items:      items,
		Pagination: meta,
	}
}

func FromPaginationResult(p *pagination.Result) *PaginationMeta {
	return &PaginationMeta{
		CurrentPage: p.CurrentPage,
		PerPage:     p.PerPage,
		Total:       p.Total,
		LastPage:    p.LastPage,
		HasNext:     p.HasNext,
		HasPrev:     p.HasPrev,
	}
}

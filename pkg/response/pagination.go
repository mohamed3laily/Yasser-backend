package response

import "yasser-backend/pkg/database"

type PaginationMeta struct {
	CurrentPage int  `json:"currentPage"`
	PerPage     int  `json:"perPage"`
	Total       int  `json:"total"`
	LastPage    int  `json:"lastPage"`
	HasNext     bool `json:"hasNext"`
	HasPrev     bool `json:"hasPrev"`
}

type PaginatedResponse struct {
	Data       interface{}     `json:"data"`
	Pagination *PaginationMeta `json:"pagination"`
}

func NewPaginatedResponse(data interface{}, meta *PaginationMeta) *PaginatedResponse {
	return &PaginatedResponse{
		Data:       data,
		Pagination: meta,
	}
}

func FromDatabasePagination(dbPagination *database.PaginationResult) *PaginationMeta {
	return &PaginationMeta{
		CurrentPage: dbPagination.CurrentPage,
		PerPage:     dbPagination.PerPage,
		Total:       int(dbPagination.Total),
		LastPage:    dbPagination.LastPage,
		HasNext:     dbPagination.HasNext,
		HasPrev:     dbPagination.HasPrev,
	}
}
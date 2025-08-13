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
    Items      interface{}     `json:"items"`
    Pagination *PaginationMeta `json:"pagination"`
}

func NewPaginatedResponse(items interface{}, meta *PaginationMeta) *PaginatedResponse {
    return &PaginatedResponse{
        Items:      items,
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
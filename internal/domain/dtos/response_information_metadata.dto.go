package dtos

type PaginationMetadataResponse struct {
	TotalItems      int  `json:"totalItems"`
	TotalPages      int  `json:"totalPages"`
	CurrentPage     int  `json:"currentPage"`
	PageSize        int  `json:"pageSize"`
	HasNextPage     bool `json:"hasNextPage"`
	HasPreviousPage bool `json:"hasPreviousPage"`
}

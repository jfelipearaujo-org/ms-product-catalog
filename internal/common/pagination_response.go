package common

type PaginationResponse[T any] struct {
	Page       int64 `json:"page"`
	TotalPages int64 `json:"total_pages"`
	TotalItems int64 `json:"total_items"`

	Data []T `json:"data"`
}

func NewPaginationResponse[T any](page, totalPages, totalItems int64, data []T) *PaginationResponse[T] {
	if data == nil {
		data = make([]T, 0)
	}

	return &PaginationResponse[T]{
		Page:       page,
		TotalPages: totalPages,
		TotalItems: totalItems,
		Data:       data,
	}
}

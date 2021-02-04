package response

import "time"

//Base is
type Base struct {
	Status     string      `json:"status"`
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Timestamp  time.Time   `json:"timestamp"`
	Data       interface{} `json:"data"`
}

// SuccessResponse is
type SuccessResponse struct {
	Base
	Data interface{} `json:"data"`
}

// SuccessReponsePagination is
type SuccessReponsePagination struct {
	SuccessResponse
	Pagination
}

// Pagination is
type Pagination struct {
	Total       int `json:"total"`
	PerPage     int `json:"per_page"`
	CurrentPage int `json:"current_page"`
	LastPage    int `json:"last_page"`
}

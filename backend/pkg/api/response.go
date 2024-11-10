package api

import "time"

type Response struct {
	Authenticated bool              `json:"authenticated"`
	User          *UserResponseInfo `json:"user,omitempty"`
	Status        string            `json:"status"`
	Message       string            `json:"message"`
	Error         *ErrorDetails     `json:"error,omitempty" swaggerignore:"true"`
	Payload       interface{}       `json:"payload,omitempty" swaggerignore:"true"`
	Pagination    interface{}       `json:"pagination,omitempty" swaggerignore:"true"`
	Metadata      Metadata          `json:"metadata"`
}

type ErrorDetails struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Details []ValidationError `json:"details,omitempty"`
}

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type GeneralPagination struct {
	TotalCount int    `json:"totalCount"`
	OrderBy    string `json:"orderBy"`
}

type Metadata struct {
	Timestamp time.Time `json:"timestamp"`
	Version   string    `json:"version"`
}

type ResponseID struct {
	ID int `json:"id"`
}

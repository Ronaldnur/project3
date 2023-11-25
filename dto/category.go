package dto

import "time"

type NewCategoryRequest struct {
	Type string `json:"type" valid:"required"`
}

type CategoryReturn struct {
	Id         int       `json:"id"`
	Type       string    `json:"type"`
	Created_at time.Time `json:"created_at"`
}

type CategoryResponse struct {
	Result     string         `json:"result"`
	StatusCode int            `json:"statusCode"`
	Message    string         `json:"message"`
	Data       CategoryReturn `json:"data"`
}

type UpdateCategoryReturn struct {
	Id         int       `json:"id"`
	Type       string    `json:"type"`
	Updated_at time.Time `json:"Updated_at"`
}
type UpdateCategoryResponse struct {
	Result     string               `json:"result"`
	StatusCode int                  `json:"statusCode"`
	Message    string               `json:"message"`
	Data       UpdateCategoryReturn `json:"data"`
}

type DeleteCategoryResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

type GetCategoryReturn struct {
	Id         int                `json:"id"`
	Type       string             `json:"type"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"Updated_at"`
	Task       GetTaskForCategory `json:"Tasks"`
}
type GetCategoryResponse struct {
	Result     string              `json:"result"`
	StatusCode int                 `json:"statusCode"`
	Message    string              `json:"message"`
	Data       []GetCategoryReturn `json:"data"`
}

package dto

import "time"

type NewTaskRequest struct {
	Title       string `json:"title" valid:"required"`
	Description string `json:"description" valid:"required"`
	Category_id int    `json:"category_id"`
}

type TaskReturn struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Status      bool      `json:"status"`
	Description string    `json:"description"`
	User_id     int       `json:"user_id"`
	Category_id int       `json:"category_id"`
	Created_at  time.Time `json:"created_at"`
}
type UpdateTaskRequest struct {
	Title       string `json:"title" valid:"required"`
	Description string `json:"description" valid:"required"`
}

type TaskResponse struct {
	Result     string     `json:"result"`
	StatusCode int        `json:"statusCode"`
	Message    string     `json:"message"`
	Data       TaskReturn `json:"data"`
}

type UpdateTaskReturn struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      bool      `json:"status"`
	User_id     int       `json:"user_id"`
	Category_id int       `json:"category_id"`
	Updated_at  time.Time `json:"updated_at"`
}

type UpdateTaskResponse struct {
	Result     string           `json:"result"`
	StatusCode int              `json:"statusCode"`
	Message    string           `json:"message"`
	Data       UpdateTaskReturn `json:"data"`
}

type UpdateStatusRequest struct {
	Status bool `json:"status" valid:"required"`
}

type UpdateCategoryIdRequest struct {
	Category_id int `json:"category_id" valid:"required"`
}

type DeleteTaskResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

type GetTaskReturn struct {
	Id          int            `json:"id"`
	Title       string         `json:"title"`
	Status      bool           `json:"status"`
	Description string         `json:"description"`
	User_id     int            `json:"user_id"`
	Category_id int            `json:"category_id"`
	Created_at  time.Time      `json:"created_at"`
	Updated_at  time.Time      `json:"updated_at"`
	User        GetUserForTask `json:"User"`
}

type GetTaskResponse struct {
	Result     string          `json:"result"`
	StatusCode int             `json:"statusCode"`
	Message    string          `json:"message"`
	Data       []GetTaskReturn `json:"data"`
}

type GetTaskForCategory struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	User_id     int       `json:"user_id"`
	Category_id int       `json:"category_id"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
}

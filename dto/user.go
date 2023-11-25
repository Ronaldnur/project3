package dto

import "time"

type NewUserRequest struct {
	Full_name string `json:"full_name" valid:"required"`
	Email     string `json:"email" valid:"email,required"`
	Password  string `json:"password" valid:"required"`
}

type UserReturn struct {
	Id         int       `json:"id"`
	Full_name  string    `json:"full_name"`
	Email      string    `json:"email"`
	Created_at time.Time `json:"created_at"`
}

type UserResponse struct {
	Result     string     `json:"result"`
	StatusCode int        `json:"statusCode"`
	Message    string     `json:"message"`
	Data       UserReturn `json:"data"`
}

type NewUserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenResponse struct {
	Token string `json:"token"`
}
type LoginResponse struct {
	Result     string        `json:"result"`
	StatusCode int           `json:"statusCode"`
	Message    string        `json:"message"`
	Data       TokenResponse `json:"token"`
}

type UpdateRequest struct {
	Full_name string `json:"full_name" valid:"required"`
	Email     string `json:"email" valid:"email,required"`
}

type UserUpdateReturn struct {
	Id         int       `json:"id"`
	Full_name  string    `json:"full_name"`
	Email      string    `json:"email"`
	Updated_at time.Time `json:"updated_at"`
}

type UpdateResponse struct {
	Result     string           `json:"result"`
	StatusCode int              `json:"statusCode"`
	Message    string           `json:"message"`
	Data       UserUpdateReturn `json:"data"`
}

type DeleteUserResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

type GetUserForTask struct {
	Id        int    `json:"id"`
	Email     string `json:"email"`
	Full_name string `json:"full_name"`
}

type PatchRoleRequest struct {
	Role string `json:"Role" valid:"required"`
}

type UserRoleReturn struct {
	Id         int       `json:"id"`
	Full_name  string    `json:"full_name"`
	Email      string    `json:"email"`
	Role       string    `json:"Role"`
	Updated_at time.Time `json:"updated_at"`
}

type UserRoleResponse struct {
	Result     string         `json:"result"`
	StatusCode int            `json:"statusCode"`
	Message    string         `json:"message"`
	Data       UserRoleReturn `json:"data"`
}

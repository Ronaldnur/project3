package service

import (
	"fmt"
	"net/http"
	"project3/dto"
	"project3/entity"
	"project3/pkg/errs"
	"project3/pkg/helpers"
	"project3/repository/user_repository"
)

type userService struct {
	userRepo user_repository.Repository
}

type UserService interface {
	CreateNewUser(payload dto.NewUserRequest) (*dto.UserResponse, errs.MessageErr)
	Login(payload dto.NewUserLogin) (*dto.LoginResponse, errs.MessageErr)
	UpdateUser(userId int, newUpdate dto.UpdateRequest) (*dto.UpdateResponse, errs.MessageErr)
	DeleteUser(userId int) (*dto.DeleteUserResponse, errs.MessageErr)
	PatchRole(userId int, payload dto.PatchRoleRequest) (*dto.UserRoleResponse, errs.MessageErr)
	SeedAdminUser() (*dto.UserResponse, errs.MessageErr)
}

func NewUserService(userRepo user_repository.Repository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}
func (u *userService) CreateNewUser(payload dto.NewUserRequest) (*dto.UserResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(payload)

	if len(payload.Password) < 6 {
		return nil, errs.NewBadRequest("Password should be at least 6 characters long")
	}

	if err != nil {
		return nil, err
	}

	existingEmail, err := u.userRepo.GetUserByEmail(payload.Email)
	if err != nil && err.Status() == http.StatusInternalServerError {
		return nil, err
	}

	if existingEmail != nil {
		return nil, errs.NewDuplicateDataError("Please Try Another Email")
	}

	user := entity.User{
		Full_name: payload.Full_name,
		Email:     payload.Email,
		Password:  payload.Password,
		Role:      "member",
	}

	err = user.HashPassword()
	if err != nil {
		return nil, err
	}

	userCreate, err := u.userRepo.CreateNewUser(user)
	if err != nil {
		return nil, err
	}

	response := dto.UserResponse{
		Result:     "success",
		StatusCode: http.StatusCreated,
		Message:    "user registered successfully",
		Data: dto.UserReturn{
			Id:         userCreate.Id,
			Full_name:  payload.Full_name,
			Email:      payload.Email,
			Created_at: userCreate.Created_at,
		},
	}

	return &response, nil
}

func (u *userService) Login(payload dto.NewUserLogin) (*dto.LoginResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(payload)

	if err != nil {
		return nil, err
	}

	user, err := u.userRepo.GetUserByEmail(payload.Email)
	if err != nil {
		if err.Status() == http.StatusNotFound {
			return nil, errs.NewBadRequest("invalid email/password")
		}
		return nil, err
	}
	isValidPassword := user.ComparePassword(payload.Password)
	if !isValidPassword {
		return nil, errs.NewBadRequest("invalid email/password")
	}

	token := user.GenerateToken()

	response := dto.LoginResponse{
		Result:     "success",
		StatusCode: http.StatusOK,
		Message:    "Login successfully",
		Data: dto.TokenResponse{
			Token: token,
		},
	}

	return &response, nil
}

func (u *userService) UpdateUser(userId int, newUpdate dto.UpdateRequest) (*dto.UpdateResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(newUpdate)
	if err != nil {
		return nil, err
	}
	existingEmail, err := u.userRepo.GetUserByEmail(newUpdate.Email)
	if err != nil && err.Status() == http.StatusInternalServerError {
		return nil, err
	}

	if existingEmail != nil {
		return nil, errs.NewDuplicateDataError("Please Try Another Email")
	}

	updateUser := entity.User{
		Full_name: newUpdate.Full_name,
		Email:     newUpdate.Email,
	}

	update, err := u.userRepo.UpdateUser(userId, updateUser)
	if err != nil {
		return nil, err
	}

	response := dto.UpdateResponse{
		Result:     "success",
		StatusCode: http.StatusOK,
		Message:    "successfully update user data",
		Data: dto.UserUpdateReturn{
			Id:         update.Id,
			Full_name:  newUpdate.Full_name,
			Email:      newUpdate.Email,
			Updated_at: update.Updated_at,
		},
	}
	return &response, nil
}

func (u *userService) DeleteUser(userId int) (*dto.DeleteUserResponse, errs.MessageErr) {

	_, err := u.userRepo.GetUserById(userId)
	if err != nil {
		return nil, err
	}

	err = u.userRepo.DeleteUser(userId)

	if err != nil {
		return nil, err
	}

	response := dto.DeleteUserResponse{
		StatusCode: http.StatusOK,
		Message:    "Your account has been successfully deleted",
	}
	return &response, nil
}

func (u *userService) PatchRole(userId int, payload dto.PatchRoleRequest) (*dto.UserRoleResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(payload)

	if err != nil {
		return nil, err
	}

	if payload.Role != "admin" {
		return nil, errs.NewBadRequest("Invalid role, It should be admin")
	}

	rolePatch := entity.User{
		Role: payload.Role,
	}

	patchRole, err := u.userRepo.PatchUserRole(userId, rolePatch)
	if err != nil {
		return nil, err
	}

	response := dto.UserRoleResponse{
		Result:     "success",
		StatusCode: http.StatusOK,
		Message:    "your role account upgraded to admin",
		Data: dto.UserRoleReturn{
			Id:         patchRole.Id,
			Full_name:  patchRole.Full_name,
			Email:      patchRole.Email,
			Role:       payload.Role,
			Updated_at: patchRole.Updated_at,
		},
	}
	return &response, nil
}

func (u *userService) SeedAdminUser() (*dto.UserResponse, errs.MessageErr) {

	existingAdmin, err := u.userRepo.GetUserByEmail("admin@gmail.com")
	if err != nil && err.Status() == http.StatusInternalServerError {
		return nil, err
	}

	if existingAdmin != nil {

		fmt.Println("Admin user already exists. Skipping admin seeding.")
		return nil, nil
	}

	adminUser := entity.User{
		Full_name: "Admin Control",
		Email:     "admin@gmail.com",
		Password:  "111111",
		Role:      "admin",
	}

	err = adminUser.HashPassword()
	if err != nil {
		return nil, err
	}

	adminUserCreate, err := u.userRepo.CreateNewUser(adminUser)
	if err != nil {
		return nil, err
	}

	response := dto.UserResponse{
		Result:     "success",
		StatusCode: http.StatusCreated,
		Message:    "Admin user registered successfully",
		Data: dto.UserReturn{
			Id:         adminUserCreate.Id,
			Full_name:  "Admin Control",
			Email:      "admin@gmail.com",
			Created_at: adminUserCreate.Created_at,
		},
	}

	return &response, nil
}

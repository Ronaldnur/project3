package user_repository

import (
	"project3/entity"
	"project3/pkg/errs"
)

type Repository interface {
	CreateNewUser(newUser entity.User) (*entity.User, errs.MessageErr)
	GetUserByEmail(userEmail string) (*entity.User, errs.MessageErr)
	UpdateUser(userId int, newUpdate entity.User) (*entity.User, errs.MessageErr)
	GetUserById(userId int) (*entity.User, errs.MessageErr)
	DeleteUser(userId int) errs.MessageErr
	PatchUserRole(userId int, role entity.User) (*entity.User, errs.MessageErr)
}

package service

import (
	"project3/entity"
	"project3/pkg/errs"
	"project3/pkg/helpers"
	"project3/repository/category_repository"
	"project3/repository/task_repository"
	"project3/repository/user_repository"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Authentitaction() gin.HandlerFunc

	AuthorizationTask() gin.HandlerFunc
	AuthorizationAdmin() gin.HandlerFunc
}

type authService struct {
	userRepo     user_repository.Repository
	categoryRepo category_repository.Repository
	taskRepo     task_repository.Repository
}

func NewAuthService(userRepo user_repository.Repository, categoryRepo category_repository.Repository, taskRepo task_repository.Repository) AuthService {
	return &authService{
		userRepo:     userRepo,
		categoryRepo: categoryRepo,
		taskRepo:     taskRepo,
	}
}
func (a *authService) Authentitaction() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var invalidTokenErr = errs.NewUnauthenticatedError("invalid token")

		bearerToken := ctx.GetHeader("Authorization")

		var user entity.User

		err := user.ValidateToken(bearerToken)

		if err != nil {
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		result, err := a.userRepo.GetUserByEmail(user.Email)

		if err != nil {
			ctx.AbortWithStatusJSON(invalidTokenErr.Status(), invalidTokenErr)
			return
		}

		_ = result
		ctx.Set("userData", user)

		ctx.Next()
	}
}

func (a *authService) AuthorizationTask() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := ctx.MustGet("userData").(entity.User)

		taskId, err := helpers.GetParamId(ctx, "taskId")

		if err != nil {
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}
		task, err := a.taskRepo.GetTaskById(taskId)

		if err != nil {
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}
		if task.User_id != user.Id {
			unauthorizedErr := errs.NewUnauthorizedError("you are not authorized to modify the task data")
			ctx.AbortWithStatusJSON(unauthorizedErr.Status(), unauthorizedErr)
			return
		}
		ctx.Next()
	}
}

func (a *authService) AuthorizationAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		user := ctx.MustGet("userData").(entity.User)
		if user.Role != "admin" {
			unauthorizedErr := errs.NewUnauthorizedError("you are not authorized to access this endpoint, only admin can access it")
			ctx.AbortWithStatusJSON(unauthorizedErr.Status(), unauthorizedErr)
			return
		}
	}
}

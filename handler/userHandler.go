package handler

import (
	"project3/dto"
	"project3/entity"
	"project3/pkg/errs"
	"project3/service"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) userHandler {
	return userHandler{
		userService: userService,
	}
}

func (uh *userHandler) Register(ctx *gin.Context) {
	var newUserRequest dto.NewUserRequest
	if err := ctx.ShouldBindJSON(&newUserRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	result, err := uh.userService.CreateNewUser(newUserRequest)
	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}
	ctx.JSON(result.StatusCode, result)
}

func (uh *userHandler) Login(ctx *gin.Context) {
	var newUserRequest dto.NewUserLogin

	if err := ctx.ShouldBindJSON(&newUserRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}
	result, err := uh.userService.Login(newUserRequest)

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}
	ctx.JSON(result.StatusCode, result)
}

func (uh *userHandler) UpdateUser(ctx *gin.Context) {
	var newUserUpdate dto.UpdateRequest

	if err := ctx.ShouldBindJSON(&newUserUpdate); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	user := ctx.MustGet("userData").(entity.User)

	result, err := uh.userService.UpdateUser(user.Id, newUserUpdate)
	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}
	ctx.JSON(result.StatusCode, result)
}

func (uh *userHandler) DeleteUser(ctx *gin.Context) {
	user := ctx.MustGet("userData").(entity.User)

	result, err := uh.userService.DeleteUser(user.Id)
	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(result.StatusCode, result)
}

func (uh *userHandler) PatchRole(ctx *gin.Context) {
	var newRole dto.PatchRoleRequest

	if err := ctx.ShouldBindJSON(&newRole); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}
	user := ctx.MustGet("userData").(entity.User)

	result, err := uh.userService.PatchRole(user.Id, newRole)
	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(result.StatusCode, result)
}

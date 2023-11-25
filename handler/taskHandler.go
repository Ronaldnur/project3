package handler

import (
	"project3/dto"
	"project3/entity"
	"project3/pkg/errs"
	"project3/pkg/helpers"
	"project3/service"

	"github.com/gin-gonic/gin"
)

type taskHandler struct {
	taskService service.TaskService
}

func NewTaskHandler(taskService service.TaskService) taskHandler {
	return taskHandler{
		taskService: taskService,
	}
}

func (th *taskHandler) CreateTask(ctx *gin.Context) {
	var newTaskRequest dto.NewTaskRequest

	if err := ctx.ShouldBindJSON(&newTaskRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}
	user := ctx.MustGet("userData").(entity.User)

	result, err := th.taskService.CreateTask(newTaskRequest, user.Id)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}
	ctx.JSON(result.StatusCode, result)
}

func (th *taskHandler) UpdateTask(ctx *gin.Context) {
	var updateRequest dto.UpdateTaskRequest
	if err := ctx.ShouldBindJSON(&updateRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	taskId, err := helpers.GetParamId(ctx, "taskId")
	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	result, err := th.taskService.UpdateTask(taskId, updateRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}
	ctx.JSON(result.StatusCode, result)
}

func (th *taskHandler) PatchStatus(ctx *gin.Context) {
	var patchRequest dto.UpdateStatusRequest
	if err := ctx.ShouldBindJSON(&patchRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}
	taskId, err := helpers.GetParamId(ctx, "taskId")
	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	result, err := th.taskService.PatchStatus(taskId, patchRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}
	ctx.JSON(result.StatusCode, result)

}

func (th *taskHandler) PatchCategoryId(ctx *gin.Context) {
	var patchRequest dto.UpdateCategoryIdRequest
	if err := ctx.ShouldBindJSON(&patchRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}
	taskId, err := helpers.GetParamId(ctx, "taskId")
	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	result, err := th.taskService.PatchCategory(taskId, patchRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}
	ctx.JSON(result.StatusCode, result)
}

func (th *taskHandler) DeleteTask(ctx *gin.Context) {
	taskId, err := helpers.GetParamId(ctx, "taskId")
	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}
	result, err := th.taskService.DeleteTask(taskId)
	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}
	ctx.JSON(result.StatusCode, result)
}

func (th *taskHandler) GetTask(ctx *gin.Context) {
	user := ctx.MustGet("userData").(entity.User)

	result, err := th.taskService.GetTaskUsers(user.Id)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}
	ctx.JSON(result.StatusCode, result)
}

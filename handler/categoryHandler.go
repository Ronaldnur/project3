package handler

import (
	"project3/dto"
	"project3/entity"
	"project3/pkg/errs"
	"project3/pkg/helpers"
	"project3/service"

	"github.com/gin-gonic/gin"
)

type categoryHandler struct {
	categoryService service.CategoryService
}

func NewCategoryHandler(categoryService service.CategoryService) categoryHandler {
	return categoryHandler{
		categoryService: categoryService,
	}
}

func (ch *categoryHandler) CreateCategory(ctx *gin.Context) {
	var newCategoryRequest dto.NewCategoryRequest
	if err := ctx.ShouldBindJSON(&newCategoryRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}
	user := ctx.MustGet("userData").(entity.User)

	result, err := ch.categoryService.CreateCategory(user.Id, newCategoryRequest)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}
	ctx.JSON(result.StatusCode, result)
}

func (ch *categoryHandler) PatchCategory(ctx *gin.Context) {
	var newCategoryRequest dto.NewCategoryRequest
	if err := ctx.ShouldBindJSON(&newCategoryRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}
	user := ctx.MustGet("userData").(entity.User)

	categoryId, err := helpers.GetParamId(ctx, "categoryId")

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	result, err := ch.categoryService.PatchCategory(user.Id, categoryId, newCategoryRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(result.StatusCode, result)
}

func (ch *categoryHandler) DeleteCategory(ctx *gin.Context) {
	user := ctx.MustGet("userData").(entity.User)

	categoryId, err := helpers.GetParamId(ctx, "categoryId")

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	result, err := ch.categoryService.DeleteCategory(user.Id, categoryId)
	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}
	ctx.JSON(result.StatusCode, result)
}

func (ch *categoryHandler) GetCategory(ctx *gin.Context) {
	user := ctx.MustGet("userData").(entity.User)

	result, err := ch.categoryService.GetCategoryData(user.Id)
	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}
	ctx.JSON(result.StatusCode, result)
}

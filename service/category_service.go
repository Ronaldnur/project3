package service

import (
	"net/http"
	"project3/dto"
	"project3/entity"
	"project3/pkg/errs"
	"project3/pkg/helpers"
	"project3/repository/category_repository"
)

type categoryService struct {
	categoryRepo category_repository.Repository
}

type CategoryService interface {
	CreateCategory(userId int, payload dto.NewCategoryRequest) (*dto.CategoryResponse, errs.MessageErr)
	PatchCategory(userId int, categoryId int, payload dto.NewCategoryRequest) (*dto.UpdateCategoryResponse, errs.MessageErr)
	DeleteCategory(userId int, categoryId int) (*dto.DeleteCategoryResponse, errs.MessageErr)
	GetCategoryData(userId int) (*dto.GetCategoryResponse, errs.MessageErr)
}

func NewCategoryService(categoryRepo category_repository.Repository) CategoryService {
	return &categoryService{
		categoryRepo: categoryRepo,
	}
}

func (c *categoryService) CreateCategory(userId int, payload dto.NewCategoryRequest) (*dto.CategoryResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(payload)

	if err != nil {
		return nil, err
	}

	newCategory := entity.Category{
		Type: payload.Type,
	}

	category, err := c.categoryRepo.CreateNewCategory(newCategory)
	if err != nil {
		return nil, err
	}
	response := dto.CategoryResponse{
		Result:     "success",
		StatusCode: http.StatusCreated,
		Message:    "category successfully created",
		Data: dto.CategoryReturn{
			Id:         category.Id,
			Type:       payload.Type,
			Created_at: category.Created_at,
		},
	}
	return &response, nil
}

func (c *categoryService) PatchCategory(userId int, categoryId int, payload dto.NewCategoryRequest) (*dto.UpdateCategoryResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(payload)
	if err != nil {
		return nil, err
	}

	categoryUpdate := entity.Category{
		Type: payload.Type,
	}

	updateCategory, err := c.categoryRepo.UpdateCategory(categoryId, categoryUpdate)
	if err != nil {
		return nil, err
	}

	response := dto.UpdateCategoryResponse{
		Result:     "success",
		StatusCode: http.StatusOK,
		Message:    "category successfully updated",
		Data: dto.UpdateCategoryReturn{
			Id:         updateCategory.Id,
			Type:       payload.Type,
			Updated_at: updateCategory.Updated_at,
		},
	}

	return &response, nil
}

func (c *categoryService) DeleteCategory(userId int, categoryId int) (*dto.DeleteCategoryResponse, errs.MessageErr) {
	_, err := c.categoryRepo.GetCategoryById(categoryId)
	if err != nil {
		return nil, err
	}

	err = c.categoryRepo.DeleteCategory(categoryId)

	if err != nil {
		return nil, err
	}

	response := dto.DeleteCategoryResponse{
		StatusCode: http.StatusOK,
		Message:    "Category has been successfully deleted",
	}

	return &response, nil
}

func (c *categoryService) GetCategoryData(userId int) (*dto.GetCategoryResponse, errs.MessageErr) {
	categories, err := c.categoryRepo.GetCategory()

	if err != nil {
		return nil, err
	}

	categoryResult := []dto.GetCategoryReturn{}

	for _, eachCategory := range categories {
		category := dto.GetCategoryReturn{
			Id:         eachCategory.Category.Id,
			Type:       eachCategory.Category.Type,
			Created_at: eachCategory.Category.Created_at,
			Updated_at: eachCategory.Category.Updated_at,
			Task:       []dto.GetTaskForCategory{},
		}

		for _, eachTask := range eachCategory.Tasks {
			task := dto.GetTaskForCategory{
				Id:          eachTask.Id,
				Title:       eachTask.Title,
				Description: eachTask.Description,
				User_id:     eachTask.User_id,
				Category_id: eachTask.Category_id,
				Created_at:  eachTask.Created_at,
				Updated_at:  eachTask.Updated_at,
			}
			category.Task = append(category.Task, task)
		}

		categoryResult = append(categoryResult, category)
	}

	response := dto.GetCategoryResponse{
		Result:     "success",
		StatusCode: http.StatusOK,
		Message:    "Get category data success",
		Data:       categoryResult,
	}
	return &response, nil
}

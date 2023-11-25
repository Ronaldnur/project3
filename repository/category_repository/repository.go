package category_repository

import (
	"project3/entity"
	"project3/pkg/errs"
)

type Repository interface {
	CreateNewCategory(newCategory entity.Category) (*entity.Category, errs.MessageErr)
	GetCategoryById(categoryId int) (*entity.Category, errs.MessageErr)
	UpdateCategory(categoryId int, update entity.Category) (*entity.Category, errs.MessageErr)
	DeleteCategory(categoryId int) errs.MessageErr
	GetCategory() ([]CategoryTaskMapped, errs.MessageErr)
}

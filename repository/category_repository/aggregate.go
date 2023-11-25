package category_repository

import "project3/entity"

type CategoryWithTask struct {
	Category entity.Category
	Task     entity.Task
}

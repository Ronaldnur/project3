package category_pg

import (
	"database/sql"
	"fmt"
	"project3/entity"
	"project3/pkg/errs"
	"project3/repository/category_repository"
)

const (
	CreateCategory = `
	INSERT INTO "category"
	(
		type
	)
	VALUES ($1)
	RETURNING id,created_at
`

	GetCategoryById = `
SELECT id,type,created_at, updated_at
FROM "category"
WHERE id = $1
`

	updateCategoryByIdQuery = `
    UPDATE "category"
    SET type = $2
    WHERE id = $1
	RETURNING id,updated_at
`

	DeleteCategoryById = `
DELETE FROM "category"
WHERE id = $1
`

	GetCategoryWithTask = `
	SELECT "c"."id", "c"."type", "c"."updated_at", "c"."created_at", "t"."id", "t"."title", "t"."description", "t"."user_id", "t"."category_id", "t"."created_at", "t"."updated_at"
	FROM "category" as "c"
	LEFT JOIN "task" as "t" ON "c"."id" = "t"."category_id"
`
)

type categoryPG struct {
	db *sql.DB
}

func NewCategoryPG(db *sql.DB) category_repository.Repository {
	return &categoryPG{
		db: db,
	}
}

func (c *categoryPG) CreateNewCategory(newCategory entity.Category) (*entity.Category, errs.MessageErr) {
	var category entity.Category

	rows := c.db.QueryRow(CreateCategory, newCategory.Type)

	err := rows.Scan(&category.Id, &category.Created_at)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("category not found")
		}
		return nil, errs.NewInternalServerError("Something went wrong")
	}
	return &category, nil
}

func (c *categoryPG) GetCategoryById(categoryId int) (*entity.Category, errs.MessageErr) {
	var category entity.Category

	rows := c.db.QueryRow(GetCategoryById, categoryId)

	err := rows.Scan(&category.Id, &category.Type, &category.Created_at, &category.Updated_at)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("category not found")
		}
		return nil, errs.NewInternalServerError("Something went wrong")
	}

	return &category, nil
}

func (c *categoryPG) UpdateCategory(categoryId int, update entity.Category) (*entity.Category, errs.MessageErr) {
	var category entity.Category
	rows := c.db.QueryRow(updateCategoryByIdQuery, categoryId, update.Type)

	err := rows.Scan(&category.Id, &category.Updated_at)
	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("category not found")
		}
		return nil, errs.NewInternalServerError("Something went wrong")
	}

	return &category, nil
}

func (c *categoryPG) DeleteCategory(categoryId int) errs.MessageErr {
	_, err := c.db.Exec(DeleteCategoryById, categoryId)
	if err != nil {
		return errs.NewInternalServerError("something went wrong")
	}
	return nil
}

func (c *categoryPG) GetCategory() ([]category_repository.CategoryTaskMapped, errs.MessageErr) {
	rows, err := c.db.Query(GetCategoryWithTask)
	if err != nil {
		return nil, errs.NewInternalServerError("something went wrong")
	}
	categoryTasks := []category_repository.CategoryWithTask{}

	for rows.Next() {
		categoryTask := categoryWithTask{}

		err = rows.Scan(
			&categoryTask.CategoryId,
			&categoryTask.CategoryType,
			&categoryTask.CategortyCreatedAt,
			&categoryTask.CategoryUpdatedAt,
			&categoryTask.TaskId,
			&categoryTask.TaskTitle,
			&categoryTask.TaskDescription,
			&categoryTask.TaskUserId,
			&categoryTask.TaskCategoryId,
			&categoryTask.TaskCreatedAt,
			&categoryTask.TaskUpdatedAt,
		)
		fmt.Println(err)
		if err != nil {
			fmt.Println(err)
			return nil, errs.NewInternalServerError("Please fill all category with task to get all category data")
		}
		categoryTasks = append(categoryTasks, *categoryTask.categoryTaskWithNull())
	}
	var result category_repository.CategoryTaskMapped

	return result.HandleMappingCategoryWithTask(categoryTasks), nil
}

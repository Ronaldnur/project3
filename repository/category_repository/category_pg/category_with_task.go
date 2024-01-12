package category_pg

import (
	"database/sql"
	"project3/entity"
	"project3/repository/category_repository"
	"time"
)

type categoryWithTask struct {
	CategoryId         int
	CategoryType       string
	CategortyCreatedAt time.Time
	CategoryUpdatedAt  time.Time
	TaskId             sql.NullInt64
	TaskTitle          sql.NullString
	TaskDescription    sql.NullString
	TaskUserId         sql.NullInt64
	TaskCategoryId     sql.NullInt64
	TaskCreatedAt      sql.NullTime
	TaskUpdatedAt      sql.NullTime
}

func (c *categoryWithTask) categoryTaskWithNull() *category_repository.CategoryWithTask {
	return &category_repository.CategoryWithTask{
		Category: entity.Category{
			Id:         c.CategoryId,
			Type:       c.CategoryType,
			Created_at: c.CategortyCreatedAt,
			Updated_at: c.CategoryUpdatedAt,
		},

		Task: entity.Task{
			Id:          int(c.TaskId.Int64),
			Title:       c.TaskTitle.String,
			Description: c.TaskDescription.String,
			User_id:     int(c.TaskUserId.Int64),
			Category_id: int(c.TaskCategoryId.Int64),
			Created_at:  c.TaskCreatedAt.Time,
			Updated_at:  c.TaskUpdatedAt.Time,
		},
	}

}

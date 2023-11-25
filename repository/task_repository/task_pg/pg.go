package task_pg

import (
	"database/sql"
	"fmt"
	"project3/entity"
	"project3/pkg/errs"
	"project3/repository/task_repository"
)

const (
	GetuserIdQuery = `
	SELECT id FROM "user" WHERE id = $1
	`
	CheckCategoryExistenceQuery = `
	SELECT id FROM "category" WHERE id = $1
	`
	CreateTaskQuery = `
	INSERT INTO "task"
	(
		title,
		description,
		status,
		user_id,
		category_id
	)
	VALUES($1, $2, $3, $4, $5)
	RETURNING id,user_id,created_at
	`

	GetTaskById = `
	SELECT id,title,description,status,user_id,category_id,created_at,updated_at
	FROM "task"
	WHERE id = $1
	`

	UpdateTask = `
	UPDATE "task"
	SET title=$2,description=$3
	WHERE id=$1
	RETURNING id,status,user_id,category_id,updated_at
	`
	PatchStatus = `
	UPDATE "task"
	SET status=$2
	WHERE id=$1
	RETURNING id,title,description,user_id,category_id,updated_at
	`

	PatchCategoryId = `
	UPDATE "task"
	SET category_id=$2
	WHERE id=$1
	RETURNING id,title,description,status,user_id,updated_at
	`
	DeleteTaskId = `
DELETE FROM "task"
WHERE id = $1
`

	GetTaskWithUser = `
SELECT "t"."id", "t"."title", "t"."status", "t"."description", "t"."user_id","t"."category_id","t"."created_at", "t"."updated_at", "u"."id", "u"."email","u"."full_name"
FROM "task" as "t"
LEFT JOIN "user" as "u" ON "t"."user_id" = "u"."id"
`
)

type taskPG struct {
	db *sql.DB
}

func NewTaskPG(db *sql.DB) task_repository.Repository {
	return &taskPG{
		db: db,
	}
}

func (t *taskPG) CreateTask(payload entity.Task, userId int) (*entity.Task, errs.MessageErr) {
	var userRow, categoryRow int
	var task entity.Task

	err := t.db.QueryRow(GetuserIdQuery, userId).Scan(&userRow)
	fmt.Println(err)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("user not found")
		}

		return nil, errs.NewInternalServerError("Something went wrong")
	}

	err = t.db.QueryRow(CheckCategoryExistenceQuery, payload.Category_id).Scan(&categoryRow)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("category not found")
		}

		return nil, errs.NewInternalServerError("Something went wrong")
	}

	err = t.db.QueryRow(CreateTaskQuery, payload.Title, payload.Description, payload.Status, userRow, payload.Category_id).Scan(&task.Id, &task.User_id, &task.Created_at)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("task not found")
		}

		return nil, errs.NewInternalServerError("Something went wrong")
	}
	return &task, nil
}

func (t *taskPG) GetTaskById(taskId int) (*entity.Task, errs.MessageErr) {
	var task entity.Task

	rows := t.db.QueryRow(GetTaskById, taskId)

	err := rows.Scan(&task.Id, &task.Title, &task.Description, &task.Status, &task.User_id, &task.Category_id, &task.Created_at, &task.Updated_at)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("task not found")
		}

		return nil, errs.NewInternalServerError("Something went wrong")
	}
	return &task, nil
}

func (t *taskPG) UpdateTask(taskId int, taskUpdate entity.Task) (*entity.Task, errs.MessageErr) {
	var task entity.Task

	rows := t.db.QueryRow(UpdateTask, taskId, taskUpdate.Title, taskUpdate.Description)

	err := rows.Scan(&task.Id, &task.Status, &task.User_id, &task.Category_id, &task.Updated_at)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("task not found")
		}

		return nil, errs.NewInternalServerError("Something went wrong")
	}
	return &task, nil
}

func (t *taskPG) PatchStatus(taskId int, taskPatch entity.Task) (*entity.Task, errs.MessageErr) {
	var task entity.Task

	rows := t.db.QueryRow(PatchStatus, taskId, taskPatch.Status)

	err := rows.Scan(&task.Id, &task.Title, &task.Description, &task.User_id, &task.Category_id, &task.Updated_at)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("task not found")
		}

		return nil, errs.NewInternalServerError("Something went wrong")
	}
	return &task, nil
}

func (t *taskPG) PatchCategory(taskId int, taskPatch entity.Task) (*entity.Task, errs.MessageErr) {
	var task entity.Task
	var categoryRow int
	baris := t.db.QueryRow(CheckCategoryExistenceQuery, taskPatch.Category_id)

	err := baris.Scan(&categoryRow)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("category not found")
		}

		return nil, errs.NewInternalServerError("Something went wrong")
	}

	rows := t.db.QueryRow(PatchCategoryId, taskId, taskPatch.Category_id)

	err = rows.Scan(&task.Id, &task.Title, &task.Description, &task.Status, &task.User_id, &task.Updated_at)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("task not found")
		}

		return nil, errs.NewInternalServerError("Something went wrong")
	}
	return &task, nil
}

func (t *taskPG) DeleteTask(taskId int) errs.MessageErr {

	_, err := t.db.Exec(DeleteTaskId, taskId)
	if err != nil {
		return errs.NewInternalServerError("something went wrong")
	}
	return nil
}

func (t *taskPG) GetTasks() (*[]task_repository.TaskWithUser, errs.MessageErr) {
	rows, err := t.db.Query(GetTaskWithUser)
	if err != nil {

		return nil, errs.NewInternalServerError("something went wrong")
	}
	taskUsers := []task_repository.TaskWithUser{}

	for rows.Next() {
		var taskUser task_repository.TaskWithUser

		err = rows.Scan(
			&taskUser.Task.Id, &taskUser.Task.Title, &taskUser.Task.Status, &taskUser.Task.Description, &taskUser.Task.User_id, &taskUser.Task.Category_id, &taskUser.Task.Created_at, &taskUser.Task.Updated_at,
			&taskUser.User.Id, &taskUser.User.Email, &taskUser.User.Full_name,
		)
		if err != nil {
			return nil, errs.NewInternalServerError("something went wrong")
		}
		taskUsers = append(taskUsers, taskUser)
	}
	return &taskUsers, nil

}

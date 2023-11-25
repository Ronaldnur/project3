package task_repository

import (
	"project3/entity"
	"project3/pkg/errs"
)

type Repository interface {
	CreateTask(payload entity.Task, userId int) (*entity.Task, errs.MessageErr)
	GetTaskById(taskId int) (*entity.Task, errs.MessageErr)
	UpdateTask(taskId int, taskUpdate entity.Task) (*entity.Task, errs.MessageErr)
	PatchStatus(taskId int, taskPatch entity.Task) (*entity.Task, errs.MessageErr)
	PatchCategory(taskId int, taskPatch entity.Task) (*entity.Task, errs.MessageErr)
	DeleteTask(taskId int) errs.MessageErr
	GetTasks() (*[]TaskWithUser, errs.MessageErr)
}

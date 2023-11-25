package task_repository

import "project3/entity"

type TaskWithUser struct {
	Task entity.Task
	User entity.User
}

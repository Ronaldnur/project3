package service

import (
	"net/http"
	"project3/dto"
	"project3/entity"
	"project3/pkg/errs"
	"project3/pkg/helpers"
	"project3/repository/task_repository"
)

type taskService struct {
	taskRepo task_repository.Repository
}

type TaskService interface {
	CreateTask(newTaskRequest dto.NewTaskRequest, userId int) (*dto.TaskResponse, errs.MessageErr)
	UpdateTask(taskId int, updateRequest dto.UpdateTaskRequest) (*dto.UpdateTaskResponse, errs.MessageErr)
	PatchStatus(taskId int, PatchRequest dto.UpdateStatusRequest) (*dto.UpdateTaskResponse, errs.MessageErr)
	PatchCategory(taskId int, PatchRequest dto.UpdateCategoryIdRequest) (*dto.UpdateTaskResponse, errs.MessageErr)
	DeleteTask(taskId int) (*dto.DeleteTaskResponse, errs.MessageErr)
	GetTaskUsers(userId int) (*dto.GetTaskResponse, errs.MessageErr)
}

func NewTaskService(taskRepo task_repository.Repository) TaskService {
	return &taskService{
		taskRepo: taskRepo,
	}
}

func (t *taskService) CreateTask(newTaskRequest dto.NewTaskRequest, userId int) (*dto.TaskResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(newTaskRequest)

	if err != nil {
		return nil, err
	}
	task := entity.Task{
		Title:       newTaskRequest.Title,
		Description: newTaskRequest.Description,
		Status:      false,
		Category_id: newTaskRequest.Category_id,
	}
	taskCreate, err := t.taskRepo.CreateTask(task, userId)

	if err != nil {
		return nil, err
	}

	response := dto.TaskResponse{
		Result:     "success",
		StatusCode: http.StatusCreated,
		Message:    "task successfully created",
		Data: dto.TaskReturn{
			Id:          taskCreate.Id,
			Title:       newTaskRequest.Title,
			Status:      taskCreate.Status,
			Description: newTaskRequest.Description,
			User_id:     taskCreate.User_id,
			Category_id: newTaskRequest.Category_id,
			Created_at:  taskCreate.Created_at,
		},
	}
	return &response, nil
}

func (t *taskService) UpdateTask(taskId int, updateRequest dto.UpdateTaskRequest) (*dto.UpdateTaskResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(updateRequest)
	if err != nil {
		return nil, err
	}

	update := entity.Task{
		Title:       updateRequest.Title,
		Description: updateRequest.Description,
	}
	taskUpdate, err := t.taskRepo.UpdateTask(taskId, update)
	if err != nil {
		return nil, err
	}

	response := dto.UpdateTaskResponse{
		Result:     "success",
		StatusCode: http.StatusOK,
		Message:    "task successfully updated",
		Data: dto.UpdateTaskReturn{
			Id:          taskUpdate.Id,
			Title:       updateRequest.Title,
			Description: updateRequest.Description,
			Status:      taskUpdate.Status,
			User_id:     taskUpdate.User_id,
			Category_id: taskUpdate.Category_id,
			Updated_at:  taskUpdate.Updated_at,
		},
	}

	return &response, nil
}

func (t *taskService) PatchStatus(taskId int, PatchRequest dto.UpdateStatusRequest) (*dto.UpdateTaskResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(PatchRequest)
	if err != nil {
		return nil, err
	}
	patch := entity.Task{
		Status: PatchRequest.Status,
	}

	taskPatch, err := t.taskRepo.PatchStatus(taskId, patch)
	if err != nil {
		return nil, err
	}

	response := dto.UpdateTaskResponse{
		Result:     "success",
		StatusCode: http.StatusOK,
		Message:    "status successfully updated",
		Data: dto.UpdateTaskReturn{
			Id:          taskPatch.Id,
			Title:       taskPatch.Title,
			Description: taskPatch.Description,
			Status:      PatchRequest.Status,
			User_id:     taskPatch.User_id,
			Category_id: taskPatch.Category_id,
			Updated_at:  taskPatch.Updated_at,
		},
	}
	return &response, nil
}

func (t *taskService) PatchCategory(taskId int, PatchRequest dto.UpdateCategoryIdRequest) (*dto.UpdateTaskResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(PatchRequest)

	if err != nil {
		return nil, err
	}

	patch := entity.Task{
		Category_id: PatchRequest.Category_id,
	}

	taskPatch, err := t.taskRepo.PatchCategory(taskId, patch)
	if err != nil {
		return nil, err
	}
	response := dto.UpdateTaskResponse{
		Result:     "success",
		StatusCode: http.StatusOK,
		Message:    "category id successfully updated",
		Data: dto.UpdateTaskReturn{
			Id:          taskPatch.Id,
			Title:       taskPatch.Title,
			Description: taskPatch.Description,
			Status:      taskPatch.Status,
			User_id:     taskPatch.User_id,
			Category_id: PatchRequest.Category_id,
			Updated_at:  taskPatch.Updated_at,
		},
	}
	return &response, nil
}

func (t *taskService) DeleteTask(taskId int) (*dto.DeleteTaskResponse, errs.MessageErr) {

	err := t.taskRepo.DeleteTask(taskId)
	if err != nil {
		return nil, err
	}

	response := dto.DeleteTaskResponse{
		StatusCode: http.StatusOK,
		Message:    "Task has been successfully deleted",
	}
	return &response, nil
}

func (t *taskService) GetTaskUsers(userId int) (*dto.GetTaskResponse, errs.MessageErr) {
	tasks, err := t.taskRepo.GetTasks()

	if err != nil {
		return nil, err
	}

	taskResult := []dto.GetTaskReturn{}

	for _, eachTask := range *tasks {
		task := dto.GetTaskReturn{
			Id:          eachTask.Task.Id,
			Title:       eachTask.Task.Title,
			Status:      eachTask.Task.Status,
			Description: eachTask.Task.Description,
			User_id:     eachTask.Task.User_id,
			Category_id: eachTask.Task.Category_id,
			Created_at:  eachTask.Task.Created_at,
			Updated_at:  eachTask.Task.Updated_at,
			User: dto.GetUserForTask{
				Id:        eachTask.User.Id,
				Email:     eachTask.User.Email,
				Full_name: eachTask.User.Full_name,
			},
		}
		taskResult = append(taskResult, task)
	}

	response := dto.GetTaskResponse{
		Result:     "Success",
		StatusCode: http.StatusOK,
		Message:    "Successfully Get All task data",
		Data:       taskResult,
	}
	return &response, nil
}

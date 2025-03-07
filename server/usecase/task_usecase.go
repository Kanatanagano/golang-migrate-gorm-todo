package usecase

import "github.com/Kanatanagano/gorm-golang-migrate-todo/entity"

type TaskUsecase interface {
	GetAllTasks() ([]entity.Task, error)
	GetTaskById(id int) (entity.Task, error)
	CreateTask(task entity.Task) (entity.Task, error)
}



package usecase

import "github.com/Kanatanagano/gorm-golang-migrate-todo/entity"

type TaskUsecase interface {
	GetAllTasks() ([]entity.Task, error)
}



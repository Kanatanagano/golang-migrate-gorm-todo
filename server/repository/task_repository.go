package repository

import "github.com/Kanatanagano/gorm-golang-migrate-todo/entity"

type Task_repository interface	{
	FindAll() ([]entity.Task, error)
	FindById(id int) (entity.Task, error)
	Create(task entity.Task) (string, error)
	Update(id int, task entity.Task) (entity.Task, error)
	Delete(id int) error
}

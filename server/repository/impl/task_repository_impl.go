package impl

import (
	"github.com/Kanatanagano/gorm-golang-migrate-todo/entity"
	"github.com/Kanatanagano/gorm-golang-migrate-todo/repository"
	"gorm.io/gorm"
)


type task_repository_impl struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) repository.Task_repository {
	return &task_repository_impl{db}
}

func (r *task_repository_impl) FindAll() ([]entity.Task, error) {
	var tasks []entity.Task
	if err := r.db.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

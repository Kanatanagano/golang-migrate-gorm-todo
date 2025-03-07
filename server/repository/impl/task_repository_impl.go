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
	return &task_repository_impl{db: db}
}

func (r *task_repository_impl) FindAll() ([]entity.Task, error) {
	var tasks []entity.Task
	if err := r.db.Preload("Labels").Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *task_repository_impl) FindById(id int) (entity.Task, error) {
	var task entity.Task
	if err := r.db.Preload("Labels").First(&task, id).Error; err != nil {
		return entity.Task{}, err
	}
	return task, nil
}

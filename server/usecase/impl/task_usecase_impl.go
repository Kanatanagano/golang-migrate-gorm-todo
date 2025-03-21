package impl

import (
	"github.com/Kanatanagano/gorm-golang-migrate-todo/entity"
	"github.com/Kanatanagano/gorm-golang-migrate-todo/repository"
	"github.com/Kanatanagano/gorm-golang-migrate-todo/usecase"
)

type taskUsecase struct {
	taskRepository repository.Task_repository
}

func NewTaskUsecase(taskRepository repository.Task_repository) usecase.TaskUsecase {
	return &taskUsecase{taskRepository: taskRepository}
}

func (u *taskUsecase) GetAllTasks() ([]entity.Task, error) {
	return u.taskRepository.FindAll()
}

func (u *taskUsecase) GetTaskById(id int) (entity.Task, error) {
	return u.taskRepository.FindById(id)
}

func (u *taskUsecase) CreateTask(task entity.Task) (string, error) {
	return u.taskRepository.Create(task)
}

func (u *taskUsecase) UpdateTask(id int, task entity.Task) (entity.Task, error) {
	return u.taskRepository.Update(id, task)
}

func (u *taskUsecase) DeleteTask(id int) error {
	return u.taskRepository.Delete(id)
}


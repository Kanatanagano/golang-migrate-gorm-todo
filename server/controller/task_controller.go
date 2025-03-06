package controller

import (
	"net/http"

	"github.com/Kanatanagano/gorm-golang-migrate-todo/usecase"
	"github.com/gin-gonic/gin"
)

type TaskController struct {
	taskUsecase usecase.TaskUsecase
}

func NewTaskController(taskUsecase usecase.TaskUsecase) *TaskController {
	return &TaskController{taskUsecase: taskUsecase}
}

func (c *TaskController) GetAllTasks(ctx *gin.Context) {
	tasks, err := c.taskUsecase.GetAllTasks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tasks)
}

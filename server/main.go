package main

import (
	"fmt"
	"log"

	"github.com/Kanatanagano/gorm-golang-migrate-todo/controller"
	"github.com/Kanatanagano/gorm-golang-migrate-todo/db"
	repository "github.com/Kanatanagano/gorm-golang-migrate-todo/repository/impl"
	usecase "github.com/Kanatanagano/gorm-golang-migrate-todo/usecase/impl"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := db.Database()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(db)
	fmt.Println("DB connected")

	taskRepository := repository.NewTaskRepository(db)
	taskUsecase := usecase.NewTaskUsecase(taskRepository)
	taskController := controller.NewTaskController(taskUsecase)
	
	router := gin.Default()
	router.GET("/tasks", taskController.GetAllTasks)
	router.Run(":8080")
}

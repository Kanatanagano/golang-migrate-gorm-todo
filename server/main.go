package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Kanatanagano/gorm-golang-migrate-todo/db"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := db.Database()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(db)
	fmt.Println("DB connected")
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
	})
	router.Run(":8080")
}

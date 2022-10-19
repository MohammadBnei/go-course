package main

import (
	"go-course/handler"
	"go-course/task"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		dbURL = "user:password@tcp(127.0.0.1:3306)/go-course?charset=utf8mb4&parseTime=True&loc=Local"
	}

	db, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	db.AutoMigrate(&task.Task{})

	taskRepository := task.NewRepository(db)
	taskService := task.NewService(taskRepository)
	taskHandler := handler.NewTaskHandler(taskService)

	r := gin.Default()
	api := r.Group("/api")
	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	api.POST("/task", taskHandler.Store)
	api.GET("/task", taskHandler.FetchAll)
	api.GET("/task/:id", taskHandler.FetchById)
	api.PUT("/task/:id", taskHandler.Update)
	api.DELETE("/task/:id", taskHandler.Delete)

	r.Run(":3000")
}

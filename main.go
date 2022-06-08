package main

import (
	"log"

	"github.com/gin-gonic/gin"
	TaskController "github.com/hugomd/go-todo/controllers"
	"github.com/hugomd/go-todo/db"
	cors "github.com/rs/cors/wrapper/gin"
)

func main() {
	log.Println("Starting server..")

	db.Init()

	r := gin.Default()
	r.Use(cors.Default())

	v1 := r.Group("/api/v1")
	{
		tasks := v1.Group("/tasks")
		{
			tasks.GET("/", TaskController.GetTasks)
			tasks.POST("/", TaskController.CreateTask)
			tasks.PUT("/:id", TaskController.UpdateTask)
			tasks.DELETE("/:id", TaskController.DeleteTask)
		}
	}

	r.Run()
}

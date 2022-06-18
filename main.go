package main

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	TaskController "github.com/hugomd/go-todo/controllers"
	"github.com/hugomd/go-todo/db"
	//cors "github.com/rs/cors/wrapper/gin"
)

func main() {
	log.Println("Starting server..")

	db.Init()

	r := gin.Default()
	//r.Use(cors.Default())

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5500"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

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

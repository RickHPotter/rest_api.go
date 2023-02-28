package main

import (
	"github.com/RickHPotter/fake_rest_api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/todos", models.GetTodos)
	router.GET("/todos/:id", models.GetTodo)

	router.POST("/todos", models.AddTodo)

	router.PATCH("/todos/:id", models.ToggleTodoStatus)

	router.DELETE("/todos/delete/", models.DeleteTodo)

	router.Run("localhost:9090")
}

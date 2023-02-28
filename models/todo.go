package models

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []Todo{
	{ID: "1", Item: "Clean Room", Completed: true},
	{ID: "2", Item: "Dirty Room", Completed: false},
	{ID: "3", Item: "Messy Room", Completed: false},
}

/*
OPERATIONS
*/

func remove(slice []Todo, index int) []Todo {
	return append(slice[:index], slice[index+1:]...)
}

/*
GET
*/

func getTodoById(id string) (*Todo, int, error) {
	for i, element := range todos {
		if element.ID == id {
			return &todos[i], i, nil
		}
	}
	return nil, -1, errors.New(NOT_FOUND)
}

func GetTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func GetTodo(context *gin.Context) {
	id := context.Param("id")
	todo, _, err := getTodoById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": NOT_FOUND})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)
}

/*
POST
*/

func AddTodo(context *gin.Context) {
	var newTodo Todo

	err := context.BindJSON(&newTodo)
	if err != nil {
		return
	}

	todos = append(todos, newTodo)

	context.IndentedJSON(http.StatusCreated, newTodo)
}

/*
PUT
*/

func ToggleTodoStatus(context *gin.Context) {
	id := context.Param("id")
	todo, _, err := getTodoById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": NOT_FOUND})
		return
	}

	todo.Completed = !todo.Completed

	context.IndentedJSON(http.StatusOK, todo)
}

/*
DELETE
*/

func DeleteTodo(context *gin.Context) {
	id, ok := context.GetQuery("id")
	if !ok {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": MISSING_ID})
		return
	}

	todo, index, err := getTodoById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": NOT_FOUND})
		return
	}

	if !todo.Completed {
		context.IndentedJSON(http.StatusNotAcceptable, gin.H{"message": UNCOMPLETED_TASK})
		return
	}

	todos = remove(todos, index)

	context.IndentedJSON(http.StatusOK, gin.H{"message": TODO_DELETED})
}

package todo

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID string `json:"id" binding:"required"`
	Item string `json:"item" binding:"required"`
	Completed bool `json:"completed" binding:"required"`
}

var todos = []todo {
	{ID: "1", Item: "Wash Plate", Completed: false},
	{ID: "2", Item: "Clean hose", Completed: false},
	{ID: "3", Item: "Jog", Completed: false},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context) {
	var newTodo todo

	err := context.ShouldBindJSON(&newTodo);

	if  err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	fmt.Println(err)

	todos = append(todos, newTodo);
	context.IndentedJSON(http.StatusCreated, todos)
}

func getTodoById(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("todo not found")
}

func getTodo(context *gin.Context) {
	var id = context.Param("id")
	todo, err := getTodoById(id)
	
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		panic("error access todo")
	}

	context.IndentedJSON(http.StatusOK, todo)
}

func updateTodo(context *gin.Context) {
	var id = context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "error updating unknown todo"})
		return
	}

	todo.Completed = !todo.Completed;

	context.IndentedJSON(http.StatusOK, todo)
}

func TodoApi() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", updateTodo)
	router.POST("/todos", addTodo)
	router.Run("localhost:9090")
}

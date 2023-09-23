package Controllers

import (
	"errors"
	"learning-golang/golang-first-api/Datas"
	"learning-golang/golang-first-api/Model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var todos = Datas.Todos

type Repository struct {
	DB *gorm.DB
}

func GetTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func AddTodo(r *Repository, context *gin.Context) {
	var newTodo Model.Todo
	// todo := Model.Todo{}

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	if err := r.DB.Create(&newTodo).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not create book"})
		return
	}

	// todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

func GetTodoById(id uint) (*Model.Todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}

	return nil, errors.New("todo not found")
}

func GetTodo(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	todo, err := GetTodoById(uint(id))
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)
}

func ToggleTodoStatus(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	todo, err := GetTodoById(uint(id))
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	todo.Completed = !todo.Completed
	context.IndentedJSON(http.StatusOK, todo)
}

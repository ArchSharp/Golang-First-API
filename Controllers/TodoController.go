package Controllers

import (
	"errors"
	"learning-golang/golang-first-api/Datas"
	"learning-golang/golang-first-api/Model"
	"net/http"

	"github.com/gin-gonic/gin"
	// uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/google/uuid"
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

func GetTodoById(id uuid.UUID) (*Model.Todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}

	return nil, errors.New("todo not found")
}

// func GetTodo(r *Repository, context *gin.Context) {
// 	newTodo := Model.Todo{}
// 	id := context.Param("id")
// 	parsedUUID := uuid.UUID{}

// 	parsedUUID.UUID.UnmarshalText([]byte(id))
// 	// todo, err := GetTodoById(parsedUUID)
// 	// if err != nil {
// 	// 	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
// 	// 	return
// 	// }

// 	// Convert the UUID to string manually
// 	uid := parsedUUID.UUID.String()

// 	if err := r.DB.Where("id = ?", uid).First(newTodo).Error; err != nil {
// 		// fmt.Println(uid)
// 		context.JSON(http.StatusBadRequest, gin.H{"message": "could not get the book"})
// 		return
// 	}

// 	context.JSON(http.StatusOK, gin.H{"message": "book id fetched successfully", "data": newTodo})

// 	// context.IndentedJSON(http.StatusOK, todo)
// }

func GetTodo(r *Repository, context *gin.Context) {
	newTodo := Model.Todo{}
	id := context.Param("id")

	// Parse the UUID from the query parameter string
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid UUID"})
		return
	}

	// Convert the UUID to string manually
	uid := parsedUUID.String()

	if err := r.DB.Where("id = ?", uid).First(&newTodo).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not get the book"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "book id fetched successfully", "data": newTodo})
}

func ToggleTodoStatus(context *gin.Context) {
	id := context.Param("id")

	// Parse the UUID from the query parameter string
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid UUID"})
		return
	}

	todo, err := GetTodoById(parsedUUID)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	todo.Completed = !todo.Completed
	context.IndentedJSON(http.StatusOK, todo)
}

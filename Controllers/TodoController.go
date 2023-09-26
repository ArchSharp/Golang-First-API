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

func GetTodos(r *Repository, context *gin.Context) {
	var todos []Model.Todo

	if err := r.DB.Find(&todos).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not get the todos"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "todos fetched successfully", "data": todos})
}

func AddTodo(r *Repository, context *gin.Context) {
	var newTodo Model.Todo

	if err := context.ShouldBindJSON(&newTodo); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := r.DB.Create(&newTodo).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not create book"})
		return
	}

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

func GetTodo(r *Repository, context *gin.Context) {
	newTodo := Model.Todo{}
	id := context.Param("id")

	// Parse the UUID from the query parameter string
	uid, err := ParseUUID(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid UUID"})
		return
	}

	if err := r.DB.Where("id = ?", uid).First(&newTodo).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not get the book"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "book id fetched successfully", "data": newTodo})
}

func PatchTodo(r *Repository, context *gin.Context) {
	id := context.Param("id")

	// Parse the UUID from the query parameter string
	uid, err := ParseUUID(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid UUID"})
		return
	}

	var todo Model.Todo

	if err := r.DB.Where("id = ?", uid).First(&todo).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	// Define a struct for parsing the JSON payload
	var updatePayload struct {
		Completed bool `json:"completed"`
	}

	// Parse the JSON request body into the updatePayload struct
	if err := context.ShouldBindJSON(&updatePayload); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON payload"})
		return
	}

	// Update the 'completed' field based on the payload
	todo.Completed = updatePayload.Completed

	if err := r.DB.Save(&todo).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update todo"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Todo updated successfully", "data": todo})
}

func UpdateTodo(r *Repository, context *gin.Context) {
	id := context.Param("id")

	// Parse the UUID from the query parameter string
	uid, err := ParseUUID(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid UUID"})
		return
	}

	var todo Model.Todo

	if err := r.DB.Where("id = ?", uid).First(&todo).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	// Define a map to hold the update fields from the payload
	var updatePayload map[string]interface{}

	// Parse the JSON request body into the updatePayload map
	if err := context.ShouldBindJSON(&updatePayload); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON payload"})
		return
	}

	// Use GORM's Updates method to update only the fields present in the payload
	if err := r.DB.Model(&todo).Updates(updatePayload).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update todo"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Todo updated successfully", "data": todo})
}

func ParseUUID(id string) (*string, error) {
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("Invalid UUID")
	}

	// Convert the UUID to string manually
	uid := parsedUUID.String()

	return &uid, nil
}

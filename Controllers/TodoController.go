package Controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"learning-golang/golang-first-api/Datas"
	"learning-golang/golang-first-api/Functions"
	"learning-golang/golang-first-api/Model"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	// "github.com/go-playground/validator/v10"

	// uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/google/uuid"
)

var todos = Datas.Todos

func GetTodos(r *Model.Repository, context *gin.Context) {
	var todos []Model.Todo

	if err := r.DB.Find(&todos).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not get the todos"})
		return
	}

	response := fmt.Sprintf("%d todos fetched successfully", len(todos))
	context.JSON(http.StatusOK, gin.H{"data": todos, "message": response})
}

func AddTodo(r *Model.Repository, context *gin.Context) {
	var newTodo Model.Todo

	if err := context.ShouldBindJSON(&newTodo); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := newTodo.Validate(); len(err) > 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
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

func GetTodo(r *Model.Repository, context *gin.Context) {
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

func PatchTodo(r *Model.Repository, context *gin.Context) {
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

func UpdateTodo(r *Model.Repository, context *gin.Context) {
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

func GetBillsCategories(context *gin.Context) {
	// Create an HTTP client with authorization headers.
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	baseURL := os.Getenv("BASEURL")
	token := os.Getenv("FLWSECK_TEST")
	client := Functions.CustomHTTPClient(token)
	var queryParam Model.GetBillsCatPayload
	if err := context.ShouldBindJSON(&queryParam); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error-1": err.Error()})
		return
	}

	urlText := fmt.Sprintf("%s?%s=%s", baseURL, queryParam.QueryParam, queryParam.Index)
	// fmt.Println(urlText)

	// You can now use the 'client' to make requests with the desired headers.
	// For example, to make a GET request:
	resp, err := client.Get(urlText)
	if err != nil {
		fmt.Printf("Error making GET request: %v\n", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to make the request"})
		return
	}
	defer resp.Body.Close()

	// Decode the response body into a generic map.
	var apiResponse map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		fmt.Printf("Error decoding response body: %v\n", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode response body"})
		return
	}

	// Set the response status and return the decoded JSON data.
	context.JSON(resp.StatusCode, apiResponse)
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

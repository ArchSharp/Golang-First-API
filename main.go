package main

import (
	"fmt"
	"learning-golang/golang-first-api/Controllers"
	"learning-golang/golang-first-api/Database"
	"learning-golang/golang-first-api/Model"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func SetupRoutes(router *gin.Engine, r *Controllers.Repository) {
	app := router.Group("/api")
	app.GET("/Todos", Controllers.GetTodos)
	app.GET("/Todos/:id", Controllers.GetTodo)
	app.PATCH("/Todos/:id", Controllers.ToggleTodoStatus)
	app.POST("/Todos", func(c *gin.Context) {
		Controllers.AddTodo(r, c) // Pass the r instance to the AddTodo function
	})
}

func main() {

	port := "5000"
	ip := "127.0.0.1"
	address := ip + ":" + port
	app := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	app.ForwardedByClientIP = true
	app.SetTrustedProxies([]string{ip})
	// app.Use(Database.Connection("user=postgres password=Alade1&&& host=localhost port=5432 dbname=quickee sslmode=disable"))
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	config := &Database.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}
	db, err := Database.NewConnection(config)

	if err != nil {
		log.Fatal("could not load the database")
	}
	err = Model.MigrateTodos(db)
	if err != nil {
		log.Fatal("could not migrate db")
	}

	r := &Controllers.Repository{
		DB: db,
	}

	SetupRoutes(app, r)
	fmt.Println("Server running on " + address)
	app.Run(address)
}

package main

import (
	"fmt"
	"learning-golang/golang-first-api/Database"
	"learning-golang/golang-first-api/Model"
	"learning-golang/golang-first-api/Routes"
	"learning-golang/golang-first-api/docs"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	port := "5000"
	ip := "127.0.0.1"
	address := ip + ":" + port
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = address
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	app := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	app.ForwardedByClientIP = true
	app.SetTrustedProxies([]string{ip})
	// CORS middleware configuration
	corsConfig := cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Change this to your actual frontend origin(s)
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"}, // Change this to your allowed headers
		AllowCredentials: true,
	})

	// Use the CORS middleware
	app.Use(corsConfig)

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

	r := &Model.Repository{
		DB: db,
	}

	Routes.TodoRoute(app, r)
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	fmt.Println("Server running on " + address)
	app.Run(address)
}

package main

import (
	"fmt"
	"learning-golang/golang-first-api/Controllers"
	"learning-golang/golang-first-api/Database"
	"learning-golang/golang-first-api/Model"
	"learning-golang/golang-first-api/Routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

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

	Routes.TodoRoute(app, r)
	fmt.Println("Server running on " + address)
	app.Run(address)
}

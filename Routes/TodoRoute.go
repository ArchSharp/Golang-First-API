package Routes

import (
	"learning-golang/golang-first-api/Controllers"
	"learning-golang/golang-first-api/Model"

	"github.com/gin-gonic/gin"
)

func TodoRoute(router *gin.Engine, r *Model.Repository) {
	app := router.Group("/api")
	app.GET("/Todos", func(c *gin.Context) {
		Controllers.GetTodos(r, c)
	})
	app.GET("/Todos/:id", func(c *gin.Context) {
		Controllers.GetTodo(r, c) // Pass the r instance to the AddTodo function
	})
	app.PATCH("/Todos/:id", func(c *gin.Context) {
		Controllers.PatchTodo(r, c)
	})
	app.PUT("/Todos/:id", func(c *gin.Context) { // this method can also be PATCH because of the way UpdateTodo is coded
		Controllers.UpdateTodo(r, c)
	})
	app.POST("/Todos", func(c *gin.Context) {
		Controllers.AddTodo(r, c) // Pass the r instance to the AddTodo function
	})

	app.POST("/BillCategories", func(c *gin.Context) {
		Controllers.GetBillsCategories(c)
	})
}

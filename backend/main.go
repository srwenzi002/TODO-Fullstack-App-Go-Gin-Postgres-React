package main

import (
	"net/http"

	api "github.com/el10savio/TODO-Fullstack-App-Go-Gin-Postgres-React/backend/api"

	"github.com/gin-gonic/gin"
)

// Function called for index
func indexView(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(http.StatusOK, gin.H{"message": "TODO APP"})
}

// Setup Gin Routes
func SetupRoutes() *gin.Engine {
	// Use Gin as router

	router := gin.Default()
	// config := cors.DefaultConfig()
	// config.AllowAllOrigins = true
	// router.Use(cors.New(config))
	// 添加统一响应中间件
	router.Use(api.CORSMiddleware())
	router.Use(api.UnifiedResponseMiddleware())

	// // Set route for index
	// router.GET("/", indexView)

	// Set routes for API
	// Update to POST, UPDATE, DELETE etc
	router.GET("/items", api.TodoItems)
	router.POST("/item/create", api.CreateTodoItem)
	router.PUT("/item/update", api.UpdateTodoItem)
	router.DELETE("/item/delete", api.DeleteTodoItem)

	// Set up Gin Server
	return router
}

// Main function
func main() {
	api.SetupPostgres()
	router := SetupRoutes()
	router.Run(":8081")
}

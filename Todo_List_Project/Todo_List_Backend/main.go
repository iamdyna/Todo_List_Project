package main

import (
	"todo-list/core"
	routers "todo-list/router"

	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var config = core.LoadConfig()

func main() {
	core.ConnectDatabase()
	// Set Gin to release mode
	gin.SetMode(gin.ReleaseMode)
	route := gin.Default()
	route.Use(cors.Default())

	todoList := route.Group("/api")
	{
		routers.TodoListRoute(todoList)
	}

	// Start the server and listen on the dedired port
	port := fmt.Sprintf(":%d", config.Port)
	fmt.Printf("Server is running on port %d", config.Port)

	err := route.Run(port)
	if err != nil {
		fmt.Println("Error starting server:", err)
		panic(err)
	}

}

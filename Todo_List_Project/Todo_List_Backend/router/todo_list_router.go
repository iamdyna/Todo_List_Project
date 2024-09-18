package routers

import (
	"todo-list/handler"
	"github.com/gin-gonic/gin"
)

func TodoListRoute(router *gin.RouterGroup){
	router.GET("/getAllTodoLists", func(ctx *gin.Context) {handler.GetAllTodoLists(ctx)} )
	router.POST("/createTodoList", func(ctx *gin.Context) {handler.CreateTodoList(ctx)})
	router.PUT("/updateTodoList/:id", func(ctx *gin.Context) {handler.UpdateTodoList(ctx)})
	router.DELETE("/deleteTodoList/:id", func(ctx *gin.Context) {handler.DeleteTodoList(ctx)})
}
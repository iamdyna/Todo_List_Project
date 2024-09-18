package handler

import (
	"fmt"
	"todo-list/core"
	"todo-list/model"
	"todo-list/utils"

	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Get all todo lists
func GetAllTodoLists(ctx *gin.Context) {

	var todoLists []model.TodoList

	err := core.Db.Find(&todoLists).Error
	if err != nil {
		fmt.Println(err.Error())
		utils.SendError(ctx, 500, "Internal server error")
		return
	}
	utils.SendData(ctx, todoLists)
}

// Create todo list
func CreateTodoList(ctx *gin.Context) {

	// Before creating todo list, you need to check the body field request first
	var body = model.TodoList{}

	// Check if there is error during biding data
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		fmt.Println(err.Error())
		utils.SendError(ctx, 400, "Invalid input")
		return
	}

	err = utils.ValidateInput(ctx, body)
	if err != nil {
		fmt.Println(err.Error())
		utils.SendError(ctx, 400, "Invalid input")
		return
	}

	// Check if the todo already exists in the database
	var existingTodo model.TodoList
	err = core.Db.Where("todo = ?", body.Todo).First(&existingTodo).Error
	if err == nil {
		utils.SendError(ctx, 400, "Todo list already exists")
		return
	}

	// Generate uuid without hyphens
	uuid := strings.Replace(uuid.New().String(), "-", "", -1)

	todoList := model.TodoList{
		ID:          uuid,
		Todo:        body.Todo,
		Iscompleted: body.Iscompleted,
	}

	err = core.Db.Create(&todoList).Error
	if err != nil {
		fmt.Println(err.Error())
		utils.SendError(ctx, 500, "Internal server error")
		return
	}

	utils.SendMessage(ctx, "Create todo list successfully")

}

// Update todo list
func UpdateTodoList(ctx *gin.Context) {

	todoListID := ctx.Param("id")
	// Check if id exist in database
	todoList := model.TodoList{}
	err := core.Db.Where("id = ?", todoListID).First(&todoList).Error
	if err != nil {
		fmt.Println(err.Error())
		utils.SendError(ctx, 404, "Todo list not found")
		return
	}

	// Before creating todo list, you need to check the body field request first
	var body = model.TodoListUpdate{}

	// Check if there is error during biding data
	err = ctx.Bind(&body)
	if err != nil {
		fmt.Println(err.Error())
		utils.SendError(ctx, 400, "Invalid input")
		return
	}

	// Validate fields
	err = utils.ValidateInput(ctx, body)
	if err != nil {
		fmt.Println(err.Error())
		utils.SendError(ctx, 400, "Invalid input")
		return
	}

	if body.Todo != "" {
		todoList.Todo = body.Todo
	}
	// Update the `is_completed` field if provided
	if body.IsCompleted != nil { // Only update if it's explicitly set
		todoList.Iscompleted = *body.IsCompleted
	}

	// Update todo list
	err = core.Db.Save(&todoList).Error
	if err != nil {
		fmt.Println(err.Error())
		utils.SendError(ctx, 500, "Internal server error")
		return
	}
	utils.SendMessage(ctx, "Update todo list successfully")
}

// Delete todo list
func DeleteTodoList(ctx *gin.Context) {
	todoListID := ctx.Param("id")
	// Check if id exist in database
	todoList := model.TodoList{}
	err := core.Db.Where("id = ?", todoListID).First(&todoList).Error
	if err != nil {
		fmt.Println(err.Error())
		utils.SendError(ctx, 404, "Todo list not found")
		return
	}

	err = core.Db.Where("id = ?", todoListID).Delete(&model.TodoList{}).Error
	if err != nil {
		fmt.Println(err.Error())
		utils.SendError(ctx, 500, "Internal server error")
		return
	}

	utils.SendMessage(ctx, "Delete todo list successfully")

}



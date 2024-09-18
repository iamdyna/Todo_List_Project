package utils

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func SendMessage(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "message": message})
}

func SendData(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{"code":200, "data":data})
}

func SendError(ctx *gin.Context, code int, message string){
	ctx.JSON(code, gin.H{"code": code, "error":message})
}
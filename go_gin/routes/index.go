package routes

import (
	"github.com/gin-gonic/gin"	
)

func Hello(ctx *gin.Context){

	params := ctx.Request.URL.Query().Get("hello")

	ctx.JSON(200, gin.H{
		"message": "pong",
		"params": params,
	})
}
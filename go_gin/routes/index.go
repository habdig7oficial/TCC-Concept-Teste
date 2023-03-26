package routes

import (
	"github.com/gin-gonic/gin"
	"go_gin/config"	
)

func Hello(ctx *gin.Context){

	dbvalue := db.Connect()

	dbquery := "hello"

	params := ctx.Request.URL.Query().Get("hello")

	ctx.JSON(200, gin.H{
		"message": "pong",
		"params": params,
		"dbvalue": dbvalue,
		"dbquery": dbquery,
	})
}
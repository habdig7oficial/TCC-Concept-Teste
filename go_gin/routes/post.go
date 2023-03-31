package routes

import (
	"fmt"
	//"go_gin/config"
	//"go_gin/config/utilities"

	"github.com/gin-gonic/gin"
)

func Post(ctx *gin.Context) {

	type dbstruct struct {
		ID     int8
		TESTE  bool
		RAMDOM float64
	}

	var newPost dbstruct

	if err := ctx.BindJSON(&newPost); err != nil {
		fmt.Print(err)
	}

	ctx.JSON(200, gin.H{
		"status": "ok",
		"Got":    newPost,
	})

}

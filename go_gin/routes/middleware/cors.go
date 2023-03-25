package cors

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func All(ctx *gin.Context){
	fmt.Print(true)
	fmt.Print(ctx)


	ctx.Header("Access-Control-Allow-Origin", "*")

	ctx.Next()
}
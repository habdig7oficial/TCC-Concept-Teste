package main 

import (
	"github.com/gin-gonic/gin"
	"go_gin/routes"
	"go_gin/routes/middleware"
)


func main(){
	GOgin := gin.Default()

	GOgin.Use(cors.All)
	
	GOgin.GET("/", routes.Hello)


	GOgin.Run("localhost:7777")
}

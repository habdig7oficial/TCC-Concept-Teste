/*
    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>
*/

package main 

import (
	"github.com/gin-gonic/gin"
	"go_gin/routes"
	"go_gin/routes/middleware"
)


func main(){
	GOgin := gin.Default()

	GOgin.Use(middleware.CORS_allow_all)
	
	GOgin.GET("/", routes.Hello)
	GOgin.POST("/", routes.Post)

	GOgin.Run("localhost:7777")
}

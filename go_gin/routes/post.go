package routes

import (
	//"go_gin/config"
	//"go_gin/config/utilities"
	//"fmt"
	"fmt"
	"log"
	"net/http"
	"go_gin/config"
	"github.com/gin-gonic/gin"
)

func Post(c *gin.Context) {

	type request struct {
		ID    int8 `form:"id"`
		Teste bool `form:"teste"`
		Random float64 `form:"random"`
		INJ string `form:"inj"`
	}

	var req request
	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Fatal(err)
	}
	

	fmt.Print(req)


	postgres := db.Connect()


	insert := postgres.QueryRow(`INSERT INTO teste(id, teste, random)
		VALUES( $1, $2, $3);`, req.ID, req.Teste, req.Random)
	fmt.Print("\n\ninsert - ", insert, "\n\n")


	var get int8
	
	mnt := fmt.Sprintf(`SELECT id FROM teste WHERE id=%s`,req.INJ)
	fmt.Print(mnt)
	selects := postgres.QueryRow(mnt).Scan(&get)

	fmt.Print("\n\nselects - ", selects, "\n\n")
	fmt.Print("\n\nget ", get, "\n\n")

	
	c.JSON(http.StatusOK, gin.H{"status": "you are logged in", "got" : req})
}
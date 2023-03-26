package routes

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"go_gin/config"	
)

func Hello(ctx *gin.Context){

	postgres := db.Connect()


	insert := postgres.QueryRow(`INSERT INTO teste(id, teste, random)
		VALUES( 2, false, 93);`)
	fmt.Print("\n\ninsert - ", insert, "\n\n")

	type dbstruct struct {
		ID     int8
		TESTE	bool
		RAMDOM	float64
	}
	

    var found []dbstruct
	id := 2

    read, err := postgres.Query("SELECT * FROM teste")
    if err != nil {
        fmt.Print(err)
		return
    }
    defer read.Close()
    // Loop through rows, using Scan to assign column data to struct fields.
    for read.Next() {
        var row dbstruct
        if err := read.Scan(&row.ID, &row.TESTE, &row.RAMDOM); err != nil {
            fmt.Print(err)
			return
        }
        found = append(found, row)
    }
    if err := read.Err(); err != nil {
        fmt.Print(err)
		return
    }
    fmt.Print(found)


	update := postgres.QueryRow(`UPDATE teste SET teste=true WHERE id=$1`, id)
	fmt.Print("\n\nupdate - ", update, "\n\n")

	delete := postgres.QueryRow(`DELETE FROM teste WHERE id=$1;`, id)
	fmt.Print("\n\ndelete - ", delete, "\n\n")

	params := ctx.Request.URL.Query().Get("hello")

	ctx.JSON(200, gin.H{
		"message": "pong",
		"params": params,
		"dbvalue": found,
		//"dbquery": dbquery,
	})
}
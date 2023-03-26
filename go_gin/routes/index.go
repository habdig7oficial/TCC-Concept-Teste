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

	license_notice := "This program is free software: you can redistribute it and/or modify\nit under the terms of the GNU General Public License as published by\nthe Free Software Foundation, either version 3 of the License, or\n(at your option) any later version.\nThis program is distributed in the hope that it will be useful,\nbut WITHOUT ANY WARRANTY; without even the implied warranty of\nMERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the\nGNU General Public License for more details.\nYou should have received a copy of the GNU General Public License\nalong with this program.  If not, see <https://www.gnu.org/licenses/>."

	ctx.JSON(200, gin.H{
		"message": "pong",
		"license": "GNU AGPL v3",
		"Free Software": "free as in freedom",
		"Author": "Mateus Felipe da Silveira Vieira",
		"COPYLEFT": "Copyright © 2023 Mateus Felipe da Silveira Vieira",
		"license_notice": license_notice,
		"query_params": params,
		"dbvalue": found,
	})
}
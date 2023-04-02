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

package config

import (
	"database/sql"
	"fmt"
)

type dbstruct struct {
	ID     int8
	TESTE  bool
	RAMDOM float64
}

func SelectAll(db *sql.DB, cols []string, where string, order_by string, IsASC bool) (found []dbstruct) {

	columns := cols

	var sort_query string
	if IsASC == true {
		sort_query = "ASC"
	} else {
		sort_query = "DESC"
	}

	fmt.Print(columns)

	var mountedQuery string
	for i := 0; i < len(columns); i++ {
		if i < len(columns)-1 {
			mountedQuery = fmt.Sprint(mountedQuery, columns[i], ", ")
		} else {
			mountedQuery = fmt.Sprint(mountedQuery, columns[i])
		}
		fmt.Print(i)
	}

	fmt.Print(mountedQuery + "\n")

	//a := fmt.Sprintf()

	read, err := db.Query("SELECT $1 FROM $2 ORDER BY $3 $4;", sql.Named("1",mountedQuery) , sql.Named("2",where), sql.Named("3",order_by), sql.Named("4",sort_query))
	if err != nil {
		fmt.Print(err)
		return found
	}
	defer read.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for read.Next() {
		var row dbstruct
		if err := read.Scan(&row.ID, &row.TESTE, &row.RAMDOM); err != nil {
			fmt.Print(err)
			return found
		}
		found = append(found, row)
	}
	if err := read.Err(); err != nil {
		fmt.Print(err)
		return found
	}
	fmt.Print(found)

	return found
}

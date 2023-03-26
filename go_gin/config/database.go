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

package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func Connect() *sql.DB{
	connStr := fmt.Sprintf("postgres://postgres:%s@db.gcuaszleoybibrbrobrp.supabase.co:6543/postgres", os.Getenv("DB_SENHA"))
	fmt.Print(connStr)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Print(err)
	}

	ping_err := db.Ping()

	if ping_err != nil {
		fmt.Print(ping_err)
	}

	return db
}
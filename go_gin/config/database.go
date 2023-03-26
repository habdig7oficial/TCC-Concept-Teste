package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func Connect() string{
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

	insert := db.QueryRow(`INSERT INTO teste(id, teste, random)
		VALUES( 2, false, 93);`)
	fmt.Print("\n\ninsert - ", insert, "\n\n")

	update := db.QueryRow(`UPDATE teste SET teste=true WHERE id=1`)
	fmt.Print("\n\nupdate - ", update, "\n\n")

	delete := db.QueryRow(`DELETE FROM teste WHERE id=2;`)
	fmt.Print("\n\ndelete - ", delete, "\n\n")


	return fmt.Sprintf("%s",db)
}
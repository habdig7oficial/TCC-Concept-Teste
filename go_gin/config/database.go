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
		VALUES( 2, true, 93)`)
	fmt.Print("\n\ninsert - ", insert, "\n\n")


	return fmt.Sprintf("%s",db)
}
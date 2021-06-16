package database

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/stdlib"
)

var (
	DB  *sql.DB
	err error
)

func Open() {
	DB, err = sql.Open("pgx", "host=database-do-user-9142180-0.b.db.ondigitalocean.com port=25060 dbname=defaultdb user=doadmin password=mmzued8avl109eeo sslmode=require")
	if err != nil {
		fmt.Println(err)
	}
}

func Close() {
	DB.Close()
}

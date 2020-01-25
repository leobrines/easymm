package sql

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/leobrines/easymm/sql/query"
	_ "github.com/lib/pq"
)

var DB *sql.DB
var Query *query.Queries

func Connect() {
	config := prodConfig()

	/*if os.Getenv("GO_ENVIRONMENT") != "production" {
		config = "file::memory:?cache=shared"
	}*/

	db, err := sql.Open("postgres", config)
	if err != nil {
		log.Fatal(err)
	}

	DB = db
	Query = query.New(db)
}

func prodConfig() string {
	host := os.Getenv("PG_HOST")
	port := os.Getenv("PG_PORT")
	user := os.Getenv("PG_USER")
	password := os.Getenv("PG_PASSWORD")
	dbname := os.Getenv("PG_DBNAME")

	connstr := "postgres://%s:%s@%s:%s/%s?sslmode=disable"

	return fmt.Sprintf(connstr, user, password, host, port, dbname)
}

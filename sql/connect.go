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
	var config string

	if os.Getenv("GO_ENVIRONMENT") == "production" {
		config = testConfig()
	} else {
		config = prodConfig()
	}

	db, err := sql.Open("postgres", config)
	if err != nil {
		log.Fatal(err)
	}

	DB = db
	Query = query.New(db)
}

func testConfig() string {
	return "file::memory:?cache=shared"
}

func prodConfig() string {
	host := os.Getenv("PG_HOST")
	port := os.Getenv("PG_PORT")
	user := os.Getenv("PG_USER")
	password := os.Getenv("PG_PASSWORD")
	dbname := os.Getenv("PG_DBNAME")

	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
}

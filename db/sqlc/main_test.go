package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var (
	dbDriver    = os.Getenv("DB_DRIVER")
	dbSource    = os.Getenv("DB_URI")
	testQueries *Queries
	testDB      *sql.DB
)

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testDB = conn
	testQueries = New(conn)

	os.Exit(m.Run())
}

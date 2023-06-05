package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver  = "postgres"
	dbSources = "postgresql://postgres:blimbeng38@localhost:5433/simple_bank?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error

	testDB, err = sql.Open(dbDriver, dbSources)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}

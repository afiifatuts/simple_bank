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

func TestMain(m *testing.M) {

	conn, err := sql.Open(dbDriver, dbSources)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}

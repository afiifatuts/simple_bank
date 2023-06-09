package main

import (
	"database/sql"
	"log"

	"github.com/afiifatuts/simple_bank/api"
	db "github.com/afiifatuts/simple_bank/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSources     = "postgresql://postgres:blimbeng38@localhost:5433/simple_bank?sslmode=disable"
	serverAddress = "localhost:8000"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSources)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}

}

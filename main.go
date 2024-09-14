package main

import (
	"database/sql"
	"log"

	"github.com/jvsena42/go_bank/api"
	db "github.com/jvsena42/go_bank/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5432/go_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
) //TODO USE ENVIROMENT VARIABLES

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)

	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}

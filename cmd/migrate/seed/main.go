package main

import (
	"log"

	"github.com/karokojnr/GoBuzz/internal/db"
	"github.com/karokojnr/GoBuzz/internal/env"
	"github.com/karokojnr/GoBuzz/internal/store"
)

func main() {
	addr := env.GetString("DB_ADDR", "postgres://admin:changeme@127.0.0.1:5432/gobuzz?sslmode=disable")

	conn, err := db.New(addr, 3, 3, "15m")
	defer conn.Close()
	if err != nil {
		log.Fatal(err)
	}

	store := store.NewStorage(conn)

	db.Seed(store, conn)
}

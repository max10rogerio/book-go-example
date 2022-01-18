package main

import (
	"github.com/max10rogerio/book-go-example/database"
	"github.com/max10rogerio/book-go-example/server"
)

func main() {
	database.StartDB()

	server := server.NewServer()

	server.Run()
}

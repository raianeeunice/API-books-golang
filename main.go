package main

import (
	"livros/database"
	"livros/server"
)

func main() {

	database.StartDB()

	server := server.NewServer()

	server.Run()
}
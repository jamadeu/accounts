package main

import (
	"github.com/jamadeu/accounts/cmd/api"
	"github.com/jamadeu/accounts/config"
)

func main() {

	db, err := config.ConnectDb()
	if err != nil {
		panic(err)
	}

	server := api.NewApiServer(":8080", db)
	if err := server.Run(); err != nil {
		panic(err)
	}
}

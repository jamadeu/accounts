package main

import (
	"os"

	"github.com/jamadeu/accounts/cmd/api"
	"github.com/jamadeu/accounts/config"
)

func main() {

	db, err := config.ConnectDb()
	if err != nil {
		panic(err)
	}

	server := api.NewApiServer(os.Getenv("PORT"), db)
	if err := server.Run(); err != nil {
		panic(err)
	}
}

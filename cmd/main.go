package main

import (
	"os"

	"github.com/jamadeu/accounts/cmd/api"
	"gorm.io/gorm"
)

func main() {

	server := api.NewApiServer(os.Getenv("PORT"), &gorm.DB{})
	if err := server.Run(); err != nil {
		panic(err)
	}
}

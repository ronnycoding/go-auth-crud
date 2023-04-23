package main

import (
	"github.com/ronnycoding/go-auth-crud/api"
)

func main() {
	router := api.SetupRouter()
	router.Run(":8000")
}

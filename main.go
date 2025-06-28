package main

import (
	"github.com/uzears/stock-suggester/routes"
)

func main() {
	router := routes.SetupRouter()
	router.Run(":8080")
}

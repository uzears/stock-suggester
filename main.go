package main

import (
	"./routes"
)

func main() {
	router := routes.SetupRouter()
	router.Run(":8080")
}

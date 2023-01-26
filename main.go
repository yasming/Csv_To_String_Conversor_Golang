package main

import (
	"csv-wrapper/infra/http/routes"
)

func main() {
	router := routes.AddRoutes()
	router.Run(":9898")
}

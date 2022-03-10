package main

import (
	"github.com/thoriqadillah/e-learning-api/config"
	"github.com/thoriqadillah/e-learning-api/routes"
)

func main() {
	config.SetupDB()
	defer config.CloseDB()

	r := routes.NewRoutes()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

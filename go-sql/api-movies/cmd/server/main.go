package main

import (
	"github.com/matiasdls1/backpack-bcgow6-matias-delaserna/go-sql/api-movies/cmd/server/routes"
	"github.com/matiasdls1/backpack-bcgow6-matias-delaserna/go-sql/api-movies/pkg/db"
)

func main() {
	engine, db := db.ConnectDatabase()
	router := routes.NewRouter(engine, db)
	router.MapRoutes()

	engine.Run(":8080")
}

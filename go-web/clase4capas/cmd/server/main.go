package main

/*
Se debe importar e inyectar el repositorio, servicio y handler
Se debe implementar el router para los diferentes endpoints
*/

import (
	"github.com/gin-gonic/gin"
	"github.com/matiasdls1/backpack-bcgow6-matias-delaserna/go-web/clase4capas/cmd/server/handler"
	"github.com/matiasdls1/backpack-bcgow6-matias-delaserna/go-web/clase4capas/internal/transactions"
)

func main() {
	repository := transactions.NewRepository()
	service := transactions.NewService(repository)
	tx := handler.NewTransaction(service)
	router := gin.Default()

	grp := router.Group("/transactions")
	grp.POST("/", tx.Store())
	grp.GET("/", tx.GetAll())
	grp.PUT("/:id", tx.Update())
	router.Run()
}

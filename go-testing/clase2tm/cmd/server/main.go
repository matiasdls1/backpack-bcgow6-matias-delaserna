package main

/*
Se debe importar e inyectar el repositorio, servicio y handler
Se debe implementar el router para los diferentes endpoints
*/

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/matiasdls1/backpack-bcgow6-matias-delaserna/go-testing/clase2tm/cmd/server/handler"
	docs "github.com/matiasdls1/backpack-bcgow6-matias-delaserna/go-testing/clase2tm/docs"
	"github.com/matiasdls1/backpack-bcgow6-matias-delaserna/go-testing/clase2tm/internal/transactions"
	"github.com/matiasdls1/backpack-bcgow6-matias-delaserna/go-testing/clase2tm/pkg/store"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle MELI Products.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	_ = godotenv.Load()
	db := store.New(store.FileType, "./transactions.json")
	repository := transactions.NewRepository(db)
	service := transactions.NewService(repository)
	tx := handler.NewTransaction(service)
	router := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	grp := router.Group("/transactions")
	grp.POST("/", tx.Store())
	grp.GET("/", tx.GetAll())
	grp.PUT("/:id", tx.Update())
	grp.DELETE("/:id", tx.Delete())
	router.Run()
}

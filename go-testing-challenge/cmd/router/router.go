package router

import (
	// "github.com/matiasdls1/bootcamp-go/desafio-cierre-testing/internal/products"
	"github.com/gin-gonic/gin"
	"github.com/matiasdls1/backpack-bcgow6-matias-delaserna/go-testing-challenge/internal/products"
)

func MapRoutes(r *gin.Engine) {
	rg := r.Group("/api/v1")
	{
		buildProductsRoutes(rg)
	}

}

func buildProductsRoutes(r *gin.RouterGroup) {
	repo := products.NewRepository()
	service := products.NewService(repo)
	handler := products.NewHandler(service)

	prodRoute := r.Group("/products")
	{
		prodRoute.GET("", handler.GetProducts)
	}

}

package main

import (
	// "github.com/bootcamp-go/desafio-cierre-testing/cmd/router"
	"github.com/gin-gonic/gin"
	"github.com/matiasdls1/backpack-bcgow6-matias-delaserna/go-testing-challenge/cmd/router"
)

func main() {
	r := gin.Default()
	router.MapRoutes(r)

	r.Run(":18085")

}

package routes

import (
	"database/sql"

	"github.com/matiasdls1/backpack-bcgow6-matias-delaserna/go-sql/api-movies/cmd/server/handler"
	"github.com/matiasdls1/backpack-bcgow6-matias-delaserna/go-sql/api-movies/internal/movie"

	"github.com/gin-gonic/gin"
)

type Router interface {
	MapRoutes()
}

type router struct {
	r  *gin.Engine
	rg *gin.RouterGroup
	db *sql.DB
}

func NewRouter(r *gin.Engine, db *sql.DB) Router {
	return &router{r: r, db: db}
}

func (r *router) MapRoutes() {
	r.setGroup()
	r.buildSellerRoutes()
}

func (r *router) setGroup() {
	r.rg = r.r.Group("/api/v1")
}

func (r *router) buildSellerRoutes() {
	repo := movie.NewRepository(r.db)
	service := movie.NewService(repo)
	handler := handler.NewMovie(service)
	r.rg.GET("/movies", handler.GetAll())
	r.rg.GET("/movies/title/:title", handler.GetByTitle())
	r.rg.GET("/movies/:id", handler.Get())
	r.rg.POST("/movies", handler.Create())
	r.rg.DELETE("/movies/:id", handler.Delete())
	r.rg.PATCH("/movies/:id", handler.Update())
}

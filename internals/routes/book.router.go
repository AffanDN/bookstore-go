package routes

import (
	"bookstore/internals/handlers"
	"bookstore/internals/middleware"
	"bookstore/internals/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func InitBookRouter(router *gin.Engine, db *sqlx.DB) {
	bookRouter := router.Group("/book")
	bookRepo := repositories.InitBookRepo(db)
	bookHandler := handlers.InitBookHandler(bookRepo)
	// localhost:8000/book
	bookRouter.GET("", bookHandler.GetBooks)
	// localhost:8000/book/new
	bookRouter.POST("/new", middleware.CheckToken, bookHandler.CreateBooks)
	// localhost:8000/book/:id
	bookRouter.GET("/:id", middleware.CheckToken, bookHandler.GetOneBook)
}

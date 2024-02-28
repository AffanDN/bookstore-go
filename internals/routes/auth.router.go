package routes

import (
	"bookstore/internals/handlers"
	"bookstore/internals/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func InitAuthRouter(router *gin.Engine, db *sqlx.DB) {
	// bikin sub router
	authRouter := router.Group("/auth")
	authRepo := repositories.InitAuthRepo(db)
	authHandler := handlers.InitAuthHandler(authRepo)
	//  Bikin rutenya
	// localhost:8000/auth/new
	authRouter.POST("/new", authHandler.Register)
	//  localhost:8000/auth
	authRouter.POST("", authHandler.Login)
}

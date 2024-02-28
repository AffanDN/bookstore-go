package handlers

import (
	"bookstore/internals/models"
	"bookstore/internals/repositories"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	*repositories.BookRepo
}

func InitBookHandler(b *repositories.BookRepo) *BookHandler {
	return &BookHandler{b}
}

func (b *BookHandler) GetBooks(ctx *gin.Context) {
	result, err := b.FindAll()
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success Get Book",
		"data":    result,
	})
}

func (b *BookHandler) CreateBooks(ctx *gin.Context) {
	body := models.BookModel{}
	if err := ctx.ShouldBind(&body); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := b.SaveBook(body); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	result, err := b.FindAll()
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Success Save Book",
		"data":    result,
	})
}

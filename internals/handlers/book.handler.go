package handlers

import (
	"bookstore/internals/models"
	"bookstore/internals/repositories"
	"log"
	"net/http"
	"strconv"

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
func (b *BookHandler) GetOneBook(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	result, err := b.FindOneBook(id)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	if len(result) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"messages": "book not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success Get 1 Book",
		"data":    result,
	})
}
func (b *BookHandler) DeleteBook(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	result, err := b.FindOneBook(id)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	if len(result) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"messages": "book not found",
		})
		return
	}
	if err := b.DeleteBookById(id); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Delete was Success",
		"data":    result,
	})
}
func (b *BookHandler) UpdateBook(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	body := models.BookModel{}
	if err := ctx.ShouldBind(&body); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	result, err := b.FindOneBook(id)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	if len(result) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"messages": "book not found",
		})
		return
	}
	if err := b.UpdateBookById(id, body); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Success Update Book",
		"data":    result,
	})
}

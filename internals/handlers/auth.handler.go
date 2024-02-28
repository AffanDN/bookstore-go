package handlers

import (
	"bookstore/internals/models"
	"bookstore/internals/repositories"
	"log"
	"net/http"

	"github.com/alexedwards/argon2id"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	*repositories.AuthRepo
}

func InitAuthHandler(a *repositories.AuthRepo) *AuthHandler {
	return &AuthHandler{a}
}

func (a *AuthHandler) Register(ctx *gin.Context) {
	// Ambil Body
	body := &models.AuthModel{}
	if err := ctx.ShouldBind(body); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return // karena ingin matiin handlernya
	}
	result, err := a.FindByEMail(*body)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	//  Pengecekan duplicate email
	if len(result) > 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Email is already registered",
		})
		return
	}

	hash, err := argon2id.CreateHash(body.Password, argon2id.DefaultParams)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := a.SaveUser(models.AuthModel{
		// Id:       body.Id,
		Email:    body.Email,
		Password: hash,
	}); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Success registered !",
	})

}
func (a *AuthHandler) Login(ctx *gin.Context) {}

package controllers

import (
	"net/http"
	"shop/user-service/configs"
	"shop/user-service/src/dtos"
	"shop/user-service/src/models"
	"shop/user-service/src/utils"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	var body dtos.RegisterDto

	if err := ctx.ShouldBindJSON(&body); err != nil {
		utils.ValidationResponse(ctx, err)
		return
	}
	db := configs.InitDB()

	user := models.User{
		Name:     body.Name,
		Email:    body.Email,
		Password: body.Password,
	}

	db.Create(&user)

	ctx.IndentedJSON(http.StatusCreated, user)
}

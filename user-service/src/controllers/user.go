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

	hash, _ := utils.HashPassword(body.Password)

	user := models.User{
		Name:     body.Name,
		Email:    body.Email,
		Password: hash,
	}

	db.Create(&user)

	ctx.IndentedJSON(http.StatusCreated, user)
}

func Login(ctx *gin.Context) {
	var body dtos.LoginDto

	if err := ctx.ShouldBindJSON(&body); err != nil {
		utils.ValidationResponse(ctx, err)
		return
	}

	//user exist check
	db := configs.InitDB()
	var user models.User
	userQuery := db.First(&user, "email=?", body.Email)
	if userQuery.RowsAffected == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	//password check
	passwordCheck := utils.CheckPasswordHash(body.Password, user.Password)
	if passwordCheck == false {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Password doesn't match"})
		return
	}

	//generate jwt token
	token, err := utils.GenerateToken(user.ID, user.Name, user.Email)
	if err != nil {
		panic(err)
	}

	ctx.IndentedJSON(http.StatusCreated, gin.H{"accessToken": token})
}

package controllers

import (
	"fmt"
	"net/http"

	"github.com/anthoturc/go/auth"
	"github.com/anthoturc/go/database"
	"github.com/anthoturc/go/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type Claims struct {
	Username string `json: "username"`
	jwt.StandardClaims
}

func RegisterUser(context *gin.Context) {

	var user models.User

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	if err := user.HashPassword(user.Password); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	data := database.Instance.Create(&user)
	datawallet := models.Wallet{Amount: 0, Amount_locked: 0, IsLocked: 0, UserId: int(user.ID)}
	walldata := database.Instance.Create(&datawallet)

	//wallet := database.Instance.Create(&wallet)
	if data.Error != nil || walldata.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": data.Error.Error()})
		context.JSON(http.StatusInternalServerError, gin.H{"error": walldata.Error.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusCreated, gin.H{"status": "Success", "code": http.StatusCreated, "data": gin.H{"userId": user.ID, "username": user.Username, "email": user.Email, "created_at": user.CreatedAt}})

}

func Signin(context *gin.Context) {
	var usr struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	type Result struct {
		Username string
		Password string
		Email    string
	}

	if err := context.ShouldBindJSON(&usr); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "incorrect parameters"})
		context.Abort()
		return
	}

	var result Result
	data := database.Instance.Table("users").Select("username", "password", "email").Where("username=?", usr.Username).Scan(&result)
	if data.Error != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("user %s not found", usr.Username)})
		context.Abort()
		return
	}

	if match := CheckPasswordHash(usr.Password, result.Password); match != true {
		context.JSON(http.StatusUnauthorized, gin.H{
			"error": "incorrect password or username",
		})
		return
	}

	token, err := auth.GenerateJWT(result.Email, result.Username)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

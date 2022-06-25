package controllers

import (
	"net/http"

	"github.com/anthoturc/go/database"
	"github.com/anthoturc/go/models"
	"github.com/gin-gonic/gin"
)

func GetProfile(context *gin.Context) {

	var user models.User
	var wallet models.Wallet
	email, ok := getSession(context)
	if !ok {
		context.JSON(http.StatusUnauthorized, gin.H{})
		return
	}
	data := database.Instance.Table("users").Select("id", "username", "email").Where("email=?", email).Scan(&user)
	selewall := database.Instance.Table("wallets").Select("*").Where("user_id=?", user.ID).Scan(&wallet)

	if data.Error != nil || selewall.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"user": data.Error.Error(), "wallet": selewall.Error.Error()}})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, gin.H{"status": "Success", "code": http.StatusOK, "data": gin.H{"user": gin.H{"userId": user.ID, "username": user.Username, "email": user.Email}, "wallet": gin.H{"wallet_id": wallet.WalletID, "amount": wallet.Amount, "amount_locked": wallet.Amount_locked}}})

}

func getSession(context *gin.Context) (string, bool) {
	email, ok := context.Get("email")
	if !ok {
		return "", false
	}

	return email.(string), true
}

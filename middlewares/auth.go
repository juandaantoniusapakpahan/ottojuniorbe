package middlewares

import (
	"net/http"
	"strings"

	"github.com/anthoturc/go/auth"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, ok := getToken(ctx)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
		}
		email, err := auth.ValidateToken(token)
		if err != nil {
			ctx.JSON(401, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}
		ctx.Writer.Header().Set("Authorization", "Bearer "+token)

		ctx.Set("email", email)

		ctx.Next()

	}
}

func getToken(contex *gin.Context) (string, bool) {
	authValue := contex.GetHeader("Authorization")
	arr := strings.Split(authValue, " ")
	if len(arr) != 2 {
		return "", false
	}

	authType := strings.Trim(arr[0], "\n\r\t")
	if strings.ToLower(authType) != strings.ToLower("Bearer") {
		return "", false
	}
	return strings.Trim(arr[1], "\n\t\r"), true
}

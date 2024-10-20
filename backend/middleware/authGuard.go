package middleware

import (
	"backend/initializer"
	"backend/models"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthGuard(ctx *gin.Context) {
	// get the cookie from the request
	tokenString, err := ctx.Cookie("Authorization")

	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

	// parse the token
	var secret string = os.Getenv("JWT_SECRET")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		log.Fatal(err)
	}

	// decode the cookie for validation

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
		var user models.User
		initializer.DB.First(&user, "email = ?", claims["email"].(string))

		if user.ID == 0 {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
		ctx.Set("user", user)
	} else {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

	// user is authenticated

	// cotinue to the next part
	ctx.Next()
}

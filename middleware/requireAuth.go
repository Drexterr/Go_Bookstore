package middleware

import (
	"net/http"
	"os"
	"time"

	"github.com/Bharat/go-bookstore/initializers"
	"github.com/Bharat/go-bookstore/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	// Get the Cookie from request
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	//validate the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {

		return []byte(os.Getenv("SECRET")), nil
	},
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	// check the exp
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		//find the user with token sub
		var User models.User
		initializers.GetDB().Where("id =?", claims["sub"]).First(&User)

		if User.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		//attach to req
		c.Set("User", User)

		//continue
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

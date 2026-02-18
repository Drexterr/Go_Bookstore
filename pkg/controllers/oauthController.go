package controllers

import (
	"net/http"

	"github.com/Bharat/go-bookstore/initializers"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
	"go.uber.org/zap"
)

func contextWithProvider(c *gin.Context) {
	provider := c.Param("provider")

	// Use Goth's official helper to safely inject the provider into the request
	c.Request = gothic.GetContextWithProvider(c.Request, provider)
}

// oauth Login function

func OauthLogin(c *gin.Context) {
	contextWithProvider(c)
	if user, err := gothic.CompleteUserAuth(c.Writer, c.Request); err == nil {
		c.JSON(http.StatusOK, gin.H{"user": user})
		return
	}
	gothic.BeginAuthHandler(c.Writer, c.Request)
}

// oauth callback function
func OauthCallback(c *gin.Context) {
	contextWithProvider(c)
	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to complete authentication"})
		return
	}

	initializers.Log.Info("User logged in via OAuth", zap.String("email: ", user.Email), zap.String("name: ", user.Name), zap.String("provider: ", user.Provider))
}

// oauth logout function
func OauthLogout(c *gin.Context) {
	contextWithProvider(c)
	gothic.Logout(c.Writer, c.Request)
	c.Redirect(http.StatusTemporaryRedirect, "/")
}

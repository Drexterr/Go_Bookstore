package middleware

import (
	"net/http"

	"github.com/Bharat/go-bookstore/pkg/models"
	"github.com/gin-gonic/gin"
)

func RequireRole(c *gin.Context) {
	// Get the user from the context

	user, exists := c.Get("User")
	if !exists {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// send value back to user structure
	userRole := user.(models.User)

	//check the role
	// if not than unauthorized
	if userRole.Role != "admin" {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	//if auth let through
	c.Next()
}

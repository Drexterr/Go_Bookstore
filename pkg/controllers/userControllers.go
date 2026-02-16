package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Bharat/go-bookstore/initializers"
	"github.com/Bharat/go-bookstore/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

// Sign UP function

func SignUp(c *gin.Context) {

	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body"})
		return
	}

	// hashing Password

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password"})
		return
	}

	// Create User

	db := initializers.GetDB()

	user := models.User{
		Email:    body.Email,
		Password: string(hash),
	}
	result := db.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user"})
		return
	}

	// respond
	c.JSON(http.StatusOK, gin.H{})
}

// Login function

func Login(c *gin.Context) {

	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body"})
		return
	}

	var User models.User
	initializers.GetDB().Where("email = ?", body.Email).First(&User)
	if User.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email or Password is incorrect"})
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(User.Password), []byte(body.Password))
	if err != nil {
		initializers.Log.Error("Email or Password is Incorrext", zap.String("email: ", body.Email))
		return

	}

	// Create JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": User.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	fmt.Println(tokenString, err)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token"})
		return
	}
	//cookie

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	// respond
	initializers.Log.Info("User logged in", zap.String("email: ", body.Email), zap.Int("userID: ", int(User.ID)))
	c.JSON(http.StatusOK, gin.H{})
}

func Validate(c *gin.Context) {

	user, _ := c.Get("User")
	c.JSON(http.StatusOK, gin.H{"I am logged in": user})
}

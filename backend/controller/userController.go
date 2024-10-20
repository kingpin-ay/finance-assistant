package controller

import (
	"backend/initializer"
	"backend/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// SignUpUser handles the registration of a new user.
func SignUpUser(ctx *gin.Context) {
	// Define a struct to capture the expected request body.
	var body struct {
		Email     string
		Password  string
		FirstName string
		LastName  string
		Age       int
	}

	// Fetch all data from the request body.
	if err := ctx.Bind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}

	// Check if the user already exists (omitted for brevity).

	// Use bcrypt to hash the password.
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to hash password"})
		return
	}

	// Create a new user object and save it to the database.
	user := models.User{
		Email:     body.Email,
		Password:  string(passwordHash),
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Age:       body.Age,
	}
	result := initializer.DB.Create(&user) // Save the user in the DB.

	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create user"})
		return
	}

	// Return a success response.
	ctx.JSON(http.StatusOK, gin.H{"message": "User Created"})
}

func Login(ctx *gin.Context) {
	var body struct {
		Email    string
		Password string
	}
	// get the email and password from the request body
	if err := ctx.Bind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}

	// check if the user exists
	var user models.User
	initializer.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password"})
		return
	}
	// check if the password is correct
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password"})
		return
	}

	// generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   user.ID,
		"firstname": user.FirstName,
		"lastname":  user.LastName,
		"email":     user.Email,
		"exp":       time.Now().Add(time.Hour * 24 * 30).UTC().Unix(),
	})

	var secret string = os.Getenv("JWT_SECRET")
	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to generate token"})
		return
	}

	// return the response as cookie
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	ctx.JSON(http.StatusOK, gin.H{})
}

func Validate(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Valid Token"})
}

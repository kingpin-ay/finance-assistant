package controller

import (
	"backend/initializer"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUpUser(ctx *gin.Context) {

	var body struct {
		Email     string
		Password  string
		FirstName string
		LastName  string
		Age       int
	}

	// fetch all data form the request body
	if ctx.Bind(&body) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
	}

	// check if the user already exists

	// user bcrypt to hash the password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to hash password"})
	}

	// save the user in the database
	var user models.User = models.User{Email: body.Email, Password: string(passwordHash), FirstName: body.FirstName, LastName: body.LastName, Age: body.Age}
	result := initializer.DB.Create(&user) // create the user in the DB.

	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create user"})
		return
	}

	// return the response
	ctx.JSON(http.StatusOK, gin.H{"message": "User Created"})
}

package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/nguyen997/gin-gorm-rest/config"
	"github.com/nguyen997/gin-gorm-rest/dto"
	"github.com/nguyen997/gin-gorm-rest/models"
	"golang.org/x/crypto/bcrypt"
)

func GetUser(ctx *gin.Context) {
	users := []models.User{}
	config.DB.Find(&users)
	ctx.JSON(200, &users)
}

func CreateUser(ctx *gin.Context) {
	var users models.User
	fmt.Println(ctx.BindJSON(&users))
	config.DB.Create(&users)
	ctx.JSON(http.StatusCreated, gin.H{"message": users})

}

func UpdateUser(ctx *gin.Context) {
	var body models.User
	// Find item by id
	if err := config.DB.Where("id = ?", ctx.Param("id")).First(&body).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	// Validate input
	var input models.User
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&body).Updates(input)
	ctx.JSON(http.StatusOK, gin.H{"data": body})
}

func DeleteUser(ctx *gin.Context) {
	var user models.User
	if err := config.DB.Where("id = ?", ctx.Param("id")).First(&user).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	config.DB.Model(&user).Delete(user)
	ctx.JSON(http.StatusOK, gin.H{"data": true})
}

func Signup(ctx *gin.Context) {
	// Get the email/password
	var body dto.CreateUserInput

	if ctx.Bind(&body) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	// Hash the password

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}
	// Create the user

	user := models.User{Email: body.Email, Password: string(hash)}
	result := config.DB.Create(&user)

	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create User",
		})
		return
	}
	// Response
	ctx.JSON(http.StatusOK, gin.H{
		"data": true,
	})
}

func Login(ctx *gin.Context) {
	// Get the email/password
	var body dto.CreateUserInput

	if ctx.Bind(&body) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	var user models.User
	config.DB.First(&user, "email = ?", body.Email)

	if user.Id == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password 1",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password 2",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, errs := token.SignedString([]byte("SECRET"))

	if errs != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid to create token",
		})
		return
	}

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", tokenString, 3600*24, "", "", false, true)
	ctx.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

func Validate(ctx *gin.Context) {
	user, _ := ctx.Get("user")

	ctx.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}

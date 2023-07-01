package controller

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"users-rest-api-gorm/dto"
	"users-rest-api-gorm/initializers"
	"users-rest-api-gorm/models"
	"users-rest-api-gorm/services"
	"users-rest-api-gorm/utils"
)

var (
	userService *services.UserService
)

func init() {
	userService = services.NewUserService()
}
func CreateUser(c *gin.Context) {
	var user models.User

	err := c.Bind(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
		return
	}

	createdUser, err := userService.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, &dto.UserDTO{ID: createdUser.ID, Name: createdUser.Name, Email: createdUser.Email, Age: createdUser.Age})
}

func GetAllUsers(c *gin.Context) {
	users, err := userService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to fetch users"})
		return
	}

	if len(users) == 0 {
		c.JSON(http.StatusOK, []models.User{})
		return
	}

	userDTOs := make([]dto.UserDTO, len(users))
	for i, user := range users {
		userDTOs[i] = dto.UserDTO{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Age:   user.Age,
		}
	}

	c.JSON(http.StatusOK, &userDTOs)
}

func GetUserById(c *gin.Context) {
	userIdStr := c.Param("id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid user id"})
		return
	}

	user, err := userService.GetUserById(userId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "no user found with such id"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to fetch user"})
		return
	}

	c.JSON(http.StatusOK, dto.UserDTO{ID: user.ID, Name: user.Name, Email: user.Email, Age: user.Age})
}

func UpdateUser(c *gin.Context) {
	userIdStr := c.Param("id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid user id"})
		return
	}

	var updatedUser dto.UpdateUserDTO

	err = c.BindJSON(&updatedUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
		return
	}

	user, err := userService.UpdateUser(userId, updatedUser)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "no user found with such id"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to update user"})
		return
	}

	c.JSON(http.StatusOK, dto.UserDTO{ID: user.ID, Name: user.Name, Email: user.Email, Age: user.Age})
}

func DeleteUser(c *gin.Context) {
	userIdStr := c.Param("id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid user id"})
		return
	}

	err = userService.DeleteUser(userId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to delete user"})
		return
	}

	c.Status(http.StatusNoContent)
}

func Login(c *gin.Context) {
	var loginRequest dto.LoginDTO

	err := c.BindJSON(&loginRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
	}

	var user models.User
	if err = initializers.DB.First(&user, "email = ?", loginRequest.Email).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
		return
	}

	//Generate JWT Token
	token, err := utils.GenerateJWTToken(user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to create token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

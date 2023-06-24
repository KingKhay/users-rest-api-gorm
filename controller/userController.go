package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"users-rest-api-gorm/initializers"
	"users-rest-api-gorm/models"
)

func CreateUser(c *gin.Context) {
	var user models.User

	err := c.Bind(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
		return
	}

	if err = initializers.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not save user"})
		return
	}

	c.JSON(http.StatusCreated, &user)
}

func GetAllUsers(c *gin.Context) {
	var users []models.User

	if err := initializers.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to fetch users. Try again later"})
		return
	}

	if len(users) == 0 {
		c.JSON(http.StatusOK, []models.User{})
		return
	}

	c.JSON(http.StatusOK, &users)
}

func GetUserById(c *gin.Context) {
	userIdStr := c.Param("id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid user id"})
		return
	}

	var user models.User
	if err = initializers.DB.First(&user, userId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		}
		return
	}

	c.JSON(http.StatusOK, &user)
}

func UpdateUser(c *gin.Context) {
	userIdStr := c.Param("id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid user id"})
		return
	}

	var updatedUser models.User

	err = c.BindJSON(&updatedUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
		return
	}

	var user models.User
	if err = initializers.DB.First(&user, userId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "no user found with such id"})
		return
	}

	initializers.DB.Model(&user).Updates(models.User{Name: updatedUser.Name, Age: updatedUser.Age, Email: updatedUser.Email})

	c.JSON(http.StatusOK, &updatedUser)
}

func DeleteUser(c *gin.Context) {
	userIdStr := c.Param("id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid user id"})
		return
	}

	var user models.User
	if err = initializers.DB.First(&user, userId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to find user"})
		return
	}

	if err = initializers.DB.Delete(&user, userId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to delete user"})
		return
	}

	c.Status(http.StatusNoContent)
}

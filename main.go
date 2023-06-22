package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"strconv"
)

type User struct {
	gorm.Model

	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func main() {
	host := os.Getenv("host")
	dbUser := os.Getenv("dbuser")
	dbName := os.Getenv("dbname")
	dbPassword := os.Getenv("dbpassword")
	dbPort := os.Getenv("dbport")
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable  port=%s", host, dbUser, dbPassword, dbName, dbPort)

	//Open Connection to database with Gorm
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	//Migrate the User struct database table
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal(err)
	}

	//Set up a gin router
	router := gin.Default()

	router.GET("/users", func(c *gin.Context) {
		var users []User
		db.Find(&users)

		if len(users) == 0 {
			c.JSON(http.StatusOK, []User{})
			return
		}

		c.JSON(http.StatusOK, &users)
	})

	router.GET("/users/:id", func(c *gin.Context) {
		userIdStr := c.Param("id")
		userId, err := strconv.Atoi(userIdStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid user id"})
			return
		}
		var user User
		if err := db.First(&user, userId).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
			}
			return
		}

		c.JSON(http.StatusOK, &user)
	})

	router.POST("/users", func(c *gin.Context) {
		var user User
		err := c.BindJSON(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
			return
		}

		db.Create(&user)
		c.JSON(http.StatusCreated, &user)
	})

	router.PUT("/users/:id", func(c *gin.Context) {
		var user User

		userIdStr := c.Param("id")
		userId, err := strconv.Atoi(userIdStr)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid user id"})
			return
		}

		var updatedUser User
		err = c.ShouldBindJSON(&updatedUser)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
			return
		}

		if err := db.First(&user, userId).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusBadRequest, gin.H{"message": "no user found"})
			}
			return
		}

		user.Name = updatedUser.Name
		user.Email = updatedUser.Email
		user.Age = updatedUser.Age

		if err = db.Save(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "could not update details"})
			return
		}
		c.JSON(http.StatusOK, &updatedUser)
	})

	router.DELETE("/users/:id", func(c *gin.Context) {
		var user User
		userIdStr := c.Param("id")
		userId, err := strconv.Atoi(userIdStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid user id"})
			return
		}

		if err = db.Find(&user, userId).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "could not find user"})
			return
		}

		if err = db.Delete(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to delete user"})
			return
		}

		c.JSON(http.StatusNoContent, gin.H{"message": "user deleted successfully"})
	})

	router.Run(":9400")

}

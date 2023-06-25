package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique" binding:"required"`
	Age      int    `json:"age"`
	Password string `json:"password" binding:"required"`
}

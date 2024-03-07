package handlers

import (
	"FIFA/models"
	"gorm.io/gorm"
	"net/mail"
)

func ValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return false
	}
	return true
}

func EmailExists(email string, db *gorm.DB) bool {
	var user models.User
	db.Where("email = ?", email).First(&user)
	if user.Email == email {
		return true
	}
	return false
}

func EmailBlacklisted(email string) bool {
	blacklist := []string{"spam@email.com"}
	for _, v := range blacklist {
		if v == email {
			return true
		}
	}
	return false
}

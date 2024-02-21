package models

import (
	"gorm.io/gorm"
)

type User struct {
	ID    uint8  `gorm:"primaryKey" json:"id"`
	Email string `json:"email"`
}

func MigrateSpots(db *gorm.DB) error {
	err := db.AutoMigrate(&User{})
	if err != nil {
		return err
	}
	return nil
}

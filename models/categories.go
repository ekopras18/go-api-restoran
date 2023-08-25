package models

import (
	"time"

	"gorm.io/gorm"
)

type Categories struct {
	ID        int       `gorm:"primaryKey;autoIncrement:true" json:"id"`
	Name      string    `gorm:"type:varchar(255)" json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// func (b *Categories) TableName() string {
// 	return "categories"
// }

func MigrateCategories(db *gorm.DB) error {
	err := db.AutoMigrate(&Categories{})
	return err
}

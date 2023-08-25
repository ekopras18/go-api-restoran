package models

import (
	"time"

	"gorm.io/gorm"
)

type Ingredients struct {
	ID        int       `gorm:"primaryKey;autoIncrement:true" json:"id"`
	Name      string    `gorm:"type:varchar(255)" json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// func (b *Ingredients) TableName() string {
// 	return "Ingredients"
// }

func MigrateIngredients(db *gorm.DB) error {
	err := db.AutoMigrate(&Ingredients{})
	return err
}

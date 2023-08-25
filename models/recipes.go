package models

import (
	"time"

	"gorm.io/gorm"
)

type Recipes struct {
	ID          int       `gorm:"primaryKey;autoIncrement:true" form:"id" json:"id"`
	Name        string    `gorm:"type:varchar(255)" form:"name" json:"name"`
	Description string    `gorm:"type:varchar(255)" form:"description" json:"description"`
	CategoryID  int       `gorm:"type:int" form:"category_id" json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	// Categories  []*Categories `gorm:"-" json:"categories"`
}

// func (b *Recipes) TableName() string {
// 	return "recipes"
// }

func MigrateRecipes(db *gorm.DB) error {
	err := db.AutoMigrate(&Recipes{})
	return err
}

package models

import (
	"time"

	"gorm.io/gorm"
)

type RecipeIngredients struct {
	ID           int       `gorm:"primaryKey;autoIncrement:true" form:"id" json:"id"`
	RecipeID     int       `gorm:"type:int" form:"recipe_id" json:"recipe_id"`
	IngredientID int       `gorm:"type:int" form:"ingredient_id" json:"ingredient_id"`
	Quantity     string    `gorm:"type:varchar(255)" form:"quantity" json:"quantity"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	// Recipes      []*Recipes     `gorm:"-" json:"recipes"`
	// Ingredients  []*Ingredients `gorm:"-" json:"ingredients"`
}

// func (b *RecipeIngredients) TableName() string {
// 	return "recipeIngredients"
// }

func MigrateRecipeIngredients(db *gorm.DB) error {
	err := db.AutoMigrate(&RecipeIngredients{})
	return err
}

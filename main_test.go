package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

// Mocking data
var (
	mockIngredient = Ingredients{
		ID:   1,
		Name: "Mock Ingredient",
	}

	mockCategory = Categories{
		ID:   1,
		Name: "Mock Category",
	}

	mockRecipe = Recipes{
		ID:          1,
		Name:        "Mock Recipe",
		Description: "Mock Description",
		CategoryID:  1,
	}

	mockRecipeIngredient = RecipeIngredients{
		ID:           1,
		RecipeID:     1,
		IngredientID: 1,
		Quantity:     "Mock Quantity",
	}
)

func TestCreateIngredient(t *testing.T) {
	app := fiber.New()
	repo := &Repository{DB: nil}
	requestBody, _ := json.Marshal(mockIngredient)
	req := httptest.NewRequest(http.MethodPost, "/api/create_ingredient", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	app.Post("/api/create_ingredient", repo.CreateIngredient)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetIngredients(t *testing.T) {
	app := fiber.New()
	repo := &Repository{DB: nil}
	requestBody, _ := json.Marshal(mockIngredient)
	req := httptest.NewRequest(http.MethodGet, "/api/get_ingredients", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	app.Get("/api/get_ingredients", repo.GetIngredient)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestUpdateIngredient(t *testing.T) {
	app := fiber.New()
	repo := &Repository{DB: nil}
	requestBody, _ := json.Marshal(mockIngredient)
	req := httptest.NewRequest(http.MethodPut, "/api/update_ingredient/1", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	app.Put("/api/update_ingredient/:id", repo.UpdateIngredient)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestCreateCategory(t *testing.T) {
	app := fiber.New()
	repo := &Repository{DB: nil}
	requestBody, _ := json.Marshal(mockCategory)
	req := httptest.NewRequest(http.MethodPost, "/api/create_category", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	app.Post("/api/create_category", repo.CreateCategory)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetCategory(t *testing.T) {
	app := fiber.New()
	repo := &Repository{DB: nil}
	requestBody, _ := json.Marshal(mockCategory)
	req := httptest.NewRequest(http.MethodGet, "/api/get_categories", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	app.Get("/api/get_categories", repo.GetCategory)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestUpdateCategory(t *testing.T) {
	app := fiber.New()
	repo := &Repository{DB: nil}
	requestBody, _ := json.Marshal(mockCategory)
	req := httptest.NewRequest(http.MethodPut, "/api/update_category/1", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	app.Put("/api/update_category/:id", repo.UpdateCategory)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestCreateRecipe(t *testing.T) {
	app := fiber.New()
	repo := &Repository{DB: nil}
	requestBody, _ := json.Marshal(mockRecipe)
	req := httptest.NewRequest(http.MethodPost, "/api/create_recipe", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	app.Post("/api/create_recipe", repo.CreateRecipe)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetRecipe(t *testing.T) {
	app := fiber.New()
	repo := &Repository{DB: nil}
	requestBody, _ := json.Marshal(mockRecipe)
	req := httptest.NewRequest(http.MethodGet, "/api/get_recipes", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	app.Get("/api/get_recipes", repo.GetRecipe)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestUpdateRecipe(t *testing.T) {
	app := fiber.New()
	repo := &Repository{DB: nil}
	requestBody, _ := json.Marshal(mockRecipe)
	req := httptest.NewRequest(http.MethodPut, "/api/update_category/1", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	app.Put("/api/update_category/:id", repo.UpdateRecipe)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestCreateRecipeingredient(t *testing.T) {
	app := fiber.New()
	repo := &Repository{DB: nil}
	requestBody, _ := json.Marshal(mockRecipeIngredient)
	req := httptest.NewRequest(http.MethodPost, "/api/create_recipeingredient", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	app.Post("/api/create_recipeingredient", repo.CreateRecipeingredient)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetRecipeingredient(t *testing.T) {
	app := fiber.New()
	repo := &Repository{DB: nil}
	requestBody, _ := json.Marshal(mockRecipeIngredient)
	req := httptest.NewRequest(http.MethodGet, "/api/get_recipeingredients", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	app.Get("/api/get_recipeingredients", repo.GetRecipeingredient)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestUpdateRecipeingredient(t *testing.T) {
	app := fiber.New()
	repo := &Repository{DB: nil}
	requestBody, _ := json.Marshal(mockRecipeIngredient)
	req := httptest.NewRequest(http.MethodPut, "/api/update_recipeingredient/1", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	app.Put("/api/update_recipeingredient/:id", repo.UpdateRecipeingredient)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestMain(m *testing.M) {
	// Run tests
	exitCode := m.Run()
	os.Exit(exitCode)
}

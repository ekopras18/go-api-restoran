package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"go-api-restoran/models"
	"go-api-restoran/storage"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Ingredients struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Categories struct {
	ID        int       `gorm:"primaryKey"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Recipes struct {
	ID          int       `gorm:"primaryKey"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CategoryID  int       `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	// Categories  []*Categories `gorm:"-" json:"categories"`
}

type RecipeIngredients struct {
	ID           int       `gorm:"primaryKey"`
	RecipeID     int       `json:"recipe_id"`
	IngredientID int       `json:"ingredient_id"`
	Quantity     string    `json:"quantity"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	// Recipes      []*Recipes `gorm:"-" json:"recipes"`
	// Ingredients  []*Ingredients `gorm:"-" json:"ingredients"`
}

type Repository struct {
	DB *gorm.DB
}

// Bahan
func (r *Repository) CreateIngredient(c *fiber.Ctx) error {
	data := Ingredients{}

	err := c.BodyParser(&data)

	if err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Cannot parse JSON"})
		return err
	}

	err = r.DB.Create(&data).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Cannot create"})
		return err
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Data has been created"})

	return nil
}

func (r *Repository) DeleteIngredient(c *fiber.Ctx) error {
	data := models.Ingredients{}
	id := c.Params("id")
	if id == "" {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Id is required"})
		return nil
	}

	res := r.DB.Delete(data, id).RowsAffected
	if res == 0 {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Cannot delete"})
		return nil
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Data has been deleted"})
	return nil
}

func (r *Repository) GetIngredient(c *fiber.Ctx) error {
	data := &[]models.Ingredients{}

	err := r.DB.Find(&data).Error

	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Cannot not get data"})
		return err
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Data has been get",
		"data":    data})

	return nil
}

func (r *Repository) GetIngredientById(c *fiber.Ctx) error {
	id := c.Params("id")
	data := &models.Ingredients{}

	if id == "" {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Id is required"})
		return nil
	}

	res := r.DB.Where("id = ?", id).First(data).RowsAffected
	if res == 0 {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Cannot not get data"})
		return nil
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Data has been get",
		"data":    data})

	return nil
}

func (r *Repository) UpdateIngredient(c *fiber.Ctx) error {
	data := models.Ingredients{}
	id := c.Params("id")

	if id == "" {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Id is required"})
		return nil
	}

	res := r.DB.First(&data, id).RowsAffected

	if res == 0 {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Cannot not get data"})
		return nil

	}

	err := c.BodyParser(&data)
	if err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Cannot parse JSON"})
		return err
	}

	// err = r.DB.Save(&data).Error
	err = r.DB.Where("id = ?", id).Updates(&data).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Cannot update"})
		return err
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Data has been updated"})
	return nil
}

// Kategori
func (r *Repository) CreateCategory(c *fiber.Ctx) error {
	data := Categories{}

	err := c.BodyParser(&data)

	if err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Cannot parse JSON"})
		return err
	}

	err = r.DB.Create(&data).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Cannot create"})
		return err
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Data has been created"})

	return nil
}

func (r *Repository) DeleteCategory(c *fiber.Ctx) error {
	data := models.Categories{}
	id := c.Params("id")
	if id == "" {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Id is required"})
		return nil
	}

	res := r.DB.Delete(data, id).RowsAffected
	if res == 0 {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Cannot delete"})
		return nil
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Data has been deleted"})
	return nil
}

func (r *Repository) GetCategory(c *fiber.Ctx) error {
	data := &[]models.Categories{}

	err := r.DB.Find(&data).Error

	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Cannot not get data"})
		return err
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Data has been get",
		"data":    data})

	return nil
}

func (r *Repository) GetCategoryById(c *fiber.Ctx) error {
	id := c.Params("id")
	data := &models.Categories{}

	if id == "" {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Id is required"})
		return nil
	}

	// fmt.Println("This Id : ", id)

	res := r.DB.Where("id = ?", id).First(data).RowsAffected
	if res == 0 {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Cannot not get data"})
		return nil
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Data has been get",
		"data":    data})

	return nil
}

func (r *Repository) UpdateCategory(c *fiber.Ctx) error {
	data := models.Categories{}
	id := c.Params("id")

	if id == "" {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Id is required"})
		return nil
	}

	res := r.DB.First(&data, id).RowsAffected

	if res == 0 {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Cannot not get data"})
		return nil

	}

	err := c.BodyParser(&data)
	if err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Cannot parse JSON"})
		return err
	}

	// err = r.DB.Save(&data).Error
	err = r.DB.Where("id = ?", id).Updates(&data).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Cannot update"})
		return err
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Data has been updated"})
	return nil
}

// Resep
func (r *Repository) CreateRecipe(c *fiber.Ctx) error {
	data := Recipes{}

	err := c.BodyParser(&data)

	if err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Cannot parse JSON"})
		return err
	}

	err = r.DB.Create(&data).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Cannot create"})
		return err
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Data has been created"})

	return nil
}

func (r *Repository) DeleteRecipe(c *fiber.Ctx) error {
	data := models.Recipes{}
	id := c.Params("id")
	if id == "" {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Id is required"})
		return nil
	}

	res := r.DB.Delete(data, id).RowsAffected
	if res == 0 {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Cannot delete"})
		return nil
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Data has been deleted"})
	return nil
}

func (r *Repository) GetRecipe(c *fiber.Ctx) error {
	data := &[]models.Recipes{}

	err := r.DB.Find(&data).Error

	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Cannot not get data"})
		return err
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Data has been get",
		"data":    data})

	return nil
}

func (r *Repository) GetRecipeById(c *fiber.Ctx) error {
	id := c.Params("id")
	data := &models.Recipes{}

	if id == "" {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Id is required"})
		return nil
	}

	// fmt.Println("This Id : ", id)

	res := r.DB.Where("id = ?", id).First(data).RowsAffected
	if res == 0 {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Cannot not get data"})
		return nil
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Data has been get",
		"data":    data})

	return nil
}

func (r *Repository) UpdateRecipe(c *fiber.Ctx) error {
	data := models.Recipes{}
	id := c.Params("id")

	if id == "" {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Id is required"})
		return nil
	}

	res := r.DB.First(&data, id).RowsAffected

	if res == 0 {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Cannot not get data"})
		return nil

	}

	err := c.BodyParser(&data)
	if err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Cannot parse JSON"})
		return err
	}

	// err = r.DB.Save(&data).Error
	err = r.DB.Where("id = ?", id).Updates(&data).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Cannot update"})
		return err
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Data has been updated"})
	return nil
}

// Resep Bahan
func (r *Repository) CreateRecipeingredient(c *fiber.Ctx) error {
	data := RecipeIngredients{}

	err := c.BodyParser(&data)

	if err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Cannot parse JSON"})
		return err
	}

	err = r.DB.Create(&data).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Cannot create"})
		return err
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Data has been created"})

	return nil
}

func (r *Repository) DeleteRecipeingredient(c *fiber.Ctx) error {
	data := models.RecipeIngredients{}
	id := c.Params("id")
	if id == "" {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Id is required"})
		return nil
	}

	res := r.DB.Delete(data, id).RowsAffected
	if res == 0 {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Cannot delete"})
		return nil
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Data has been deleted"})
	return nil
}

func (r *Repository) GetRecipeingredient(c *fiber.Ctx) error {
	data := &[]models.RecipeIngredients{}

	err := r.DB.Find(&data).Error

	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Cannot not get data"})
		return err
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Data has been get",
		"data":    data})

	return nil
}

func (r *Repository) GetRecipeingredientById(c *fiber.Ctx) error {
	id := c.Params("id")
	data := &models.RecipeIngredients{}

	if id == "" {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Id is required"})
		return nil
	}

	// fmt.Println("This Id : ", id)

	res := r.DB.Where("id = ?", id).First(data).RowsAffected
	if res == 0 {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Cannot not get data"})
		return nil
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Data has been get",
		"data":    data})

	return nil
}

func (r *Repository) UpdateRecipeingredient(c *fiber.Ctx) error {
	data := models.RecipeIngredients{}
	id := c.Params("id")

	if id == "" {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Id is required"})
		return nil
	}

	res := r.DB.First(&data, id).RowsAffected

	if res == 0 {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Cannot not get data"})
		return nil

	}

	err := c.BodyParser(&data)
	if err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Cannot parse JSON"})
		return err
	}

	// err = r.DB.Save(&data).Error
	err = r.DB.Where("id = ?", id).Updates(&data).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Cannot update"})
		return err
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Data has been updated"})
	return nil
}

func (r *Repository) FilterRecipeingredient(c *fiber.Ctx) error {
	data := &[]models.RecipeIngredients{}
	recipe := c.Query("recipe")
	ingredient := c.Query("ingredient")

	if recipe == "" && ingredient == "" {
		res := r.DB.Find(&data).RowsAffected
		if res == 0 {
			c.Status(http.StatusNotFound).JSON(
				&fiber.Map{"message": "Data not found"})
			return nil
		}

		c.Status(http.StatusOK).JSON(&fiber.Map{
			"message": "Data has been get",
			"data":    data})
	}

	if recipe != "" && ingredient == "" {
		res := r.DB.Where("recipe_id = ?", recipe).Find(&data).RowsAffected
		if res == 0 {
			c.Status(http.StatusNotFound).JSON(
				&fiber.Map{"message": "Data not found"})
			return nil
		}

		c.Status(http.StatusOK).JSON(&fiber.Map{
			"message": "Data has been get",
			"data":    data})
	}

	if recipe == "" && ingredient != "" {
		res := r.DB.Where("ingredient_id = ?", ingredient).Find(&data).RowsAffected
		if res == 0 {
			c.Status(http.StatusNotFound).JSON(
				&fiber.Map{"message": "Data not found"})
			return nil
		}

		c.Status(http.StatusOK).JSON(&fiber.Map{
			"message": "Data has been get",
			"data":    data})
	}

	if recipe != "" && ingredient != "" {
		res := r.DB.Where("recipe_id = ? AND ingredient_id = ?", recipe, ingredient).Find(&data).RowsAffected
		if res == 0 {
			c.Status(http.StatusNotFound).JSON(
				&fiber.Map{"message": "Data not found"})
			return nil
		}

		c.Status(http.StatusOK).JSON(&fiber.Map{
			"message": "Data has been get",
			"data":    data})
	}
	return nil
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	//Bahan
	api.Post("/create_ingredient", r.CreateIngredient)
	api.Delete("/delete_ingredient/:id", r.DeleteIngredient)
	api.Get("/get_ingredients", r.GetIngredient)
	api.Get("/get_ingredient/:id", r.GetIngredientById)
	api.Put("/update_ingredient/:id", r.UpdateIngredient)

	//Kategori
	api.Post("/create_category", r.CreateCategory)
	api.Delete("/delete_category/:id", r.DeleteCategory)
	api.Get("/get_categories", r.GetCategory)
	api.Get("/get_category/:id", r.GetCategoryById)
	api.Put("/update_category/:id", r.UpdateCategory)

	//Resep
	api.Post("/create_recipe", r.CreateRecipe)
	api.Delete("/delete_recipe/:id", r.DeleteRecipe)
	api.Get("/get_recipes", r.GetRecipe)
	api.Get("/get_recipe/:id", r.GetRecipeById)
	api.Put("/update_recipe/:id", r.UpdateRecipe)

	//Resep Bahan
	api.Post("/create_recipeingredient", r.CreateRecipeingredient)
	api.Delete("/delete_recipeingredient/:id", r.DeleteRecipeingredient)
	api.Get("/get_recipeingredients", r.GetRecipeingredient)
	api.Get("/get_recipeingredient/:id", r.GetRecipeingredientById)
	api.Put("/update_recipeingredient/:id", r.UpdateRecipeingredient)
	api.Get("/filter_recipeingredient", r.FilterRecipeingredient)

}

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	config := storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASSWORD"),
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	db, err := storage.NewConnection(config)

	if err != nil {
		log.Fatal("Could not connect to the database", err)
	}

	// Migrate the schema
	// Ingredient
	err = models.MigrateIngredients(db)
	if err != nil {
		log.Fatal("Could not migrate Ingredients", err)
	}
	// Category
	err = models.MigrateCategories(db)
	if err != nil {
		log.Fatal("Could not migrate Categories", err)
	}

	// Recipes
	err = models.MigrateRecipes(db)
	if err != nil {
		log.Fatal("Could not migrate Recipes", err)
	}

	// RecipeIngredients
	err = models.MigrateRecipeIngredients(db)
	if err != nil {
		log.Fatal("Could not migrate RecipeIngredients", err)
	}

	r := Repository{
		DB: db,
	}
	app := fiber.New()
	r.SetupRoutes(app)
	log.Println("Server started on: http://localhost:8080")
	app.Listen(":8080")

}

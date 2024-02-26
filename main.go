package main

import (
	"FIFA/handlers"
	"FIFA/models"
	"FIFA/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

type Users struct {
	Email string `json:"email"`
}

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) RegisterEmail(c *fiber.Ctx) error {
	UserModel := Users{}
	err := c.BodyParser(&UserModel)

	if err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message": "Failed to parse user"})
		return err
	}

	err = r.DB.Create(UserModel).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "failed to create database entry"})
		return err
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{"message": "user registered successfully!"})
	return nil
}

func (r *Repository) GetUsers(c *fiber.Ctx) error {
	UserModels := &[]models.User{}
	err := r.DB.Find(&UserModels).Error

	if err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message": "failed to find user"})
		return err
	}
	c.Status(http.StatusOK).JSON(&fiber.Map{"message": "users found successfully!", "data": UserModels})
	return nil
}

func (r *Repository) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	userModel := &models.User{}

	if id == "" {
		c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message": "empty ID"})
		return nil
	}

	err := r.DB.Where("id = ?", id).Delete(&userModel).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "failed to delete user"})
		return err
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{"message": "user deleted successfully"})
	return nil

}

func (r *Repository) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	userModel := &models.User{}
	newModel := Users{}

	err := c.BodyParser(&newModel)
	if err != nil {
		return err
	}

	if id == "" {
		c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message": "empty ID"})
		return nil
	}
	err = r.DB.Model(userModel).Where("id = ?", id).Updates(newModel).Error

	if err != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "failed to update user"})
		return err
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{"message": "user updated successfully", "data": userModel})
	return nil
}

func (r *Repository) GetUsersEmails() ([]string, error) {
	var emails []string
	var users []models.User
	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	for _, user := range users {
		emails = append(emails, user.Email)
	}
	return emails, nil
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/FIFA")
	api.Get("/users", r.GetUsers)
	api.Post("/create", r.RegisterEmail)
	api.Delete("/delete", r.DeleteUser)
	api.Put("/update", r.UpdateUser)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Panic(err)
	}

	//emails := os.Getenv("TO")
	//emailAddresses := strings.Split(emails, ",")

	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	r := Repository{
		DB: db,
	}

	err = models.MigrateSpots(db)
	if err != nil {
		log.Fatal("Failed to migrate database")
	}

	emailAddresses, err := r.GetUsersEmails()

	err = handlers.SendMail(emailAddresses)
	if err != nil {
		log.Panic(err)
	}

	app := fiber.New()
	r.SetupRoutes(app)

	err = app.Listen(":8080")
	if err != nil {
		log.Fatal("Failed to listen on port 8080")
	}

	//for {
	//	err = handlers.SendMail(emailAddresses)
	//	if err != nil {
	//		log.Panic(err)
	//	}
	//}

}

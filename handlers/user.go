package handlers

import (
	"boilerplate/database"
	"boilerplate/models"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

// UserList func gets all existing users
//
//	@Description	Get all existing users
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	models.User
//	@router			/users [get]
func UserList(c *fiber.Ctx) error {
	db := database.GetDB()

	var users []*models.User
	if err := db.Limit(10).Find(&users).Error; err != nil {
		log.Printf("Error fetching users: %v\n", err)
		return nil
	}

	return c.JSON(fiber.Map{
		"success": true,
		"users":   users,
	})
}

// UserList func creates a new user
//
//	@Description	Create a new user
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			users	body		models.CreateUser	true	"Create User"
//	@Success		200		{object}	models.User
//	@router			/users [post]
func UserCreate(c *fiber.Ctx) error {
	db := database.GetDB()

	createUser := new(models.CreateUser)

	if err := c.BodyParser(createUser); err != nil {
		log.Printf("Error parsing JSON: %v\n", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "Cannot parse JSON",
		})
	}

	if err := validate.Struct(&createUser); err != nil {
		log.Printf("Validation error: %v\n", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	user := &models.User{
		Username: createUser.Username,
		Email:    createUser.Email,
		Password: createUser.Password,
	}

	if err := db.Create(user).Error; err != nil {
		log.Printf("Error inserting user: %v\n", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "Cannot insert user",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"user":    user,
	})
}

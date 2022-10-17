package routes

import (
	"github.com/cumhurmesuttokalak/go-fiberApp2/database"
	"github.com/cumhurmesuttokalak/go-fiberApp2/helpers"
	"github.com/cumhurmesuttokalak/go-fiberApp2/models"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func CreateResponseUser(userModel models.User) User {
	return User{ID: userModel.ID, FirstName: userModel.FirstName, LastName: userModel.LastName}
}
func CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		c.Status(503).SendString("Error creating user")
		return nil
	}
	database.Database.Db.Create(&user)
	responseUser := CreateResponseUser(user)
	err := c.JSON(&responseUser)
	helpers.CheckError(err)
	return nil
}
func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	database.Database.Db.Find(&users)
	responseUsers := []User{}
	for _, user := range users {
		responseUser := CreateResponseUser(user)
		responseUsers = append(responseUsers, responseUser)
	}
	return c.Status(200).JSON(&responseUsers)
}
func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user User
	database.Database.Db.Find(&user, id)
	c.JSON(user)
	return nil
}
func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user User
	database.Database.Db.First(&user, id)
	if user.FirstName == " " {
		c.Status(500).SendString("No user found with given ID")
		return nil
	}
	database.Database.Db.Delete(&user, id)
	return nil
}

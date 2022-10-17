package main

import (
	"github.com/cumhurmesuttokalak/go-fiberApp2/database"
	"github.com/cumhurmesuttokalak/go-fiberApp2/routes"
	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to the fiberApp2")
}
func Setuproutes(app *fiber.App) error {
	app.Get("/", welcome)
	app.Get("/api/getusers", routes.GetUsers)
	app.Get("/api/getuser/:id", routes.GetUser)
	app.Post("/api/createuser", routes.CreateUser)
	app.Delete("/api/deleteuser/:id", routes.DeleteUser)
	return nil
}
func main() {
	database.ConnectDb()
	app := fiber.New()
	Setuproutes(app)
	app.Listen(":4000")
}

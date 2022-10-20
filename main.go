package main

import (
	"github.com/cumhurmesuttokalak/go-fiberApp2/database"
	"github.com/cumhurmesuttokalak/go-fiberApp2/routes"
	"github.com/gofiber/fiber/v2"
)

func Setuproutes(app *fiber.App) {
	//user endpoints

	app.Get("/api/getusers", routes.GetUsers)
	app.Get("/api/getuser/:id", routes.GetUser)
	app.Post("/api/createuser", routes.CreateUser)
	app.Delete("/api/deleteuser/:id", routes.DeleteUser)
	app.Put("/api/updateuser/:id", routes.UpdateUser)

	//product endpoints

	app.Get("/api/getproducts", routes.GetProducts)
	app.Get("/api/getproduct/:id", routes.GetProduct)
	app.Post("/api/createproduct", routes.CreateProduct)
	app.Delete("/api/deleteproduct/:id", routes.DeleteProduct)
	app.Put("/api/updateproduct/:id", routes.UpdateProduct)

	//order endpoints
	app.Get("/api/getorders", routes.GetOrders)
	app.Get("/api/getorder/:id", routes.GetOrder)
	app.Post("/api/createorder", routes.CreateOrder)
}
func main() {
	database.ConnectDb()
	app := fiber.New()
	Setuproutes(app)
	app.Listen(":4000")
}

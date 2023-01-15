package controllers

import (
	"errors"
	"time"

	"github.com/cumhurmesuttokalak/go-fiberApp2/database"
	"github.com/cumhurmesuttokalak/go-fiberApp2/models"
	"github.com/gofiber/fiber/v2"
)

type Order struct {
	ID        uint      `json:"id"`
	User      User      `json:"user"`
	Product   Product   `json:"product"`
	CreatedAt time.Time `json:"created_at"`
}

func CreateResponseOrder(order models.Order, user User, product Product) Order {
	return Order{
		ID:        order.ID,
		User:      user,
		Product:   product,
		CreatedAt: order.CreatedAt,
	}
}
func findOrder(id int, order *models.Order) error {
	database.Database.Db.Find(&order, "id=?", id)
	if order.ID == 0 {
		return errors.New("order does not exist")
	}
	return nil
}
func GetOrders(c *fiber.Ctx) error {
	var orders []models.Order
	database.Database.Db.Find(&orders)
	var responseOrders = []Order{}
	for _, order := range orders {
		var user models.User
		var product models.Product
		database.Database.Db.Find(&user, "id=?", order.UserRefer)
		database.Database.Db.Find(&product, "id=?", order.ProductRefer)
		responseOrder := CreateResponseOrder(order, CreateResponseUser(user), CreateResponseProduct(product))
		responseOrders = append(responseOrders, responseOrder)
	}
	return c.Status(200).JSON(responseOrders)
}
func GetOrder(c *fiber.Ctx) error {
	var order models.Order
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}
	if err := findOrder(id, &order); err != nil {
		c.Status(400).JSON(err.Error())
	}
	var user models.User
	var product models.Product
	database.Database.Db.First(&user, order.UserRefer)
	database.Database.Db.First(&product, order.Product)
	responseUser := CreateResponseUser(user)
	responseProduct := CreateResponseProduct(product)
	responseOrder := CreateResponseOrder(order, responseUser, responseProduct)
	return c.Status(200).JSON(responseOrder)

}
func CreateOrder(c *fiber.Ctx) error {
	var order models.Order
	err := c.BodyParser(&order)
	if err != nil {
		return err
	}
	var user models.User
	if err := findUser(order.UserRefer, &user); err != nil {
		return err
	}
	var product models.Product
	if err := findProduct(order.ProductRefer, &product); err != nil {
		return err
	}
	database.Database.Db.Create(&order)
	responseUser := CreateResponseUser(user)
	responseProduct := CreateResponseProduct(product)
	responseOrder := CreateResponseOrder(order, responseUser, responseProduct)
	return c.Status(503).JSON(responseOrder)
}

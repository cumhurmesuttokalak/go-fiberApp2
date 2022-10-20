package routes

import (
	"errors"

	"github.com/cumhurmesuttokalak/go-fiberApp2/database"
	"github.com/cumhurmesuttokalak/go-fiberApp2/models"
	"github.com/gofiber/fiber/v2"
)

type Product struct {
	ID            uint   `json:"ID"`
	Name          string `json:"name"`
	Serial_number string `json:"serial_number"`
}

func CreateResponseProduct(productModel models.Product) Product {
	return Product{ID: productModel.ID, Name: productModel.Name, Serial_number: productModel.SerialNumber}
}
func findProduct(id int, product *models.Product) error {
	database.Database.Db.Find(&product, "id=?", id)
	if product.ID == 0 {
		return errors.New("product does not exist")
	}
	return nil
}
func GetProducts(c *fiber.Ctx) error {
	var products []models.Product
	var responseProducts []Product
	database.Database.Db.Find(&products)
	for _, product := range products {
		responseProduct := CreateResponseProduct(product)
		responseProducts = append(responseProducts, responseProduct)
	}
	return c.Status(503).JSON(&responseProducts)
}
func GetProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("Please ensure that:id is an integer")
	}
	var product models.Product
	if err := findProduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	responseProduct := CreateResponseProduct(product)
	return c.Status(400).JSON(responseProduct)
}
func CreateProduct(c *fiber.Ctx) error {
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		c.Status(503).SendString("Error creating product: ")
	}
	database.Database.Db.Create(&product)
	responseProduct := CreateResponseProduct(product)
	return c.Status(200).JSON(responseProduct)
}
func DeleteProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var product models.Product
	if err != nil {
		return c.Status(350).JSON("Please ensure that :id is an integer")
	}
	if err := findProduct(id, &product); err != nil {
		return c.Status(404).JSON(err.Error())
	}
	if err := database.Database.Db.Delete(&product).Error; err != nil {
		return c.Status(500).JSON(err.Error())
	}
	return c.Status(200).SendString("Successfully deleted Product")
}
func UpdateProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		c.Status(503).JSON("Please ensure that :id is an integer")
	}
	var product models.Product
	if err := findProduct(id, &product); err != nil {
		c.Status(503).JSON(err.Error())
	}
	type Updateproduct struct {
		Name         string `json:"name"`
		Serialnumber string `json:"serial_number"`
	}
	var updateData Updateproduct
	if err := c.BodyParser(&updateData); err != nil {
		c.Status(503).JSON("Error updating product")
	}
	product.Name = updateData.Name
	product.SerialNumber = updateData.Serialnumber
	database.Database.Db.Save(product)
	responseProduct := CreateResponseProduct(product)
	return c.Status(200).JSON(responseProduct)
}

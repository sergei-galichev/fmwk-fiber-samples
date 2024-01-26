package server

import (
	"encoding/json"
	"express-style/internal/domain/product"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/pkg/errors"
)

func (s *server) CreateProduct(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")
	p := new(product.Product)
	err := c.BodyParser(p)
	if err != nil {
		log.Error("Error parsing body to structure")
		return c.Status(fiber.StatusBadRequest).JSON(
			&fiber.Map{
				"success": false,
				"message": err.Error(),
			},
		)
	}
	err = s.productService.CreateProduct(p)
	if err != nil {
		log.Error(err)
		return c.Status(fiber.StatusInternalServerError).JSON(
			&fiber.Map{
				"success": false,
				"message": err.Error(),
			},
		)
	} else {
		log.Debug("Product created: ", p)
		return c.Status(fiber.StatusCreated).JSON(
			&fiber.Map{
				"success": true,
				"message": "Product created",
			},
		)
	}
}

func (s *server) GetAllProducts(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")
	prods, err := s.productService.GetAllProducts()
	if err != nil {
		log.Error(err)
		return c.Status(fiber.StatusInternalServerError).JSON(
			&fiber.Map{
				"success": false,
				"message": err.Error(),
			},
		)
	}
	bytes, err := json.Marshal(prods)
	if err != nil {
		log.Error("Error json marshalling")
		return c.Status(fiber.StatusInternalServerError).JSON(
			&fiber.Map{
				"success": false,
				"message": errors.New("Error json marshalling"),
			},
		)
	}

	_, err = c.Write(bytes)
	if err != nil {
		log.Error("Error writing data to body")
		return c.Status(fiber.StatusInternalServerError).JSON(
			&fiber.Map{
				"success": false,
				"message": errors.New("Error writing data body"),
			},
		)
	} else {
		log.Debug("All products: ", &prods)
		return nil
	}
}

func (s *server) GetSingleProduct(c *fiber.Ctx) error {
	return nil
}

func (s *server) DeleteProduct(c *fiber.Ctx) error {
	return nil
}

func (s *server) Info(c *fiber.Ctx) error {
	_, err := c.WriteString("Info page")
	if err != nil {
		log.Error("Error writing string")
	}

	return nil
}

func (s *server) PrintAllParams(c *fiber.Ctx) error {
	bytes, err := json.MarshalIndent(c.AllParams(), "", " ")
	if err != nil {
		return errors.New("Error marshaling json")
	}
	_, err = c.Write(bytes)
	if err != nil {
		return errors.New("Error writing string")
	}
	return nil
}

func (s *server) GetSettings(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")
	return c.JSON(c.App().Stack())
}

func (s *server) Bind(c *fiber.Ctx) error {
	return nil
}

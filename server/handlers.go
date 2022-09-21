package server

import (
	"github.com/gofiber/fiber/v2"
)

func IndexHandler(c *fiber.Ctx) error {
	return c.SendString("Hello")
}
func PostHandler(c *fiber.Ctx) error {
	return c.SendString("Hello")
}
func PutHandler(c *fiber.Ctx) error {
	return c.SendString("Hello")
}
func DeleteHandler(c *fiber.Ctx) error {
	return c.SendString("Hello")
}

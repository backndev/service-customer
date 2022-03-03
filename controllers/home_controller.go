package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {
	fmt.Fprintf(c, "Welcome to Training Sharing Vision Indonesia!")

	return nil
}

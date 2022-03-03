package controllers

import (
	"github.com/gofiber/fiber/v2"
	"service-customer/database"
	"service-customer/models"
	"service-customer/util"
	"strconv"
	"time"
)

func Login(c *fiber.Ctx) error {
	var db = database.DB

	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return err
	}

	db.Where("Name = ?", user.Name).First(&user)

	if user.Id == 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "name not found",
		})
	}

	if err := user.ComparePassword(user.Password); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	token, err := util.GenerateJwt(strconv.Itoa(int(user.Id)))

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

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
	var data map[string]string
	db := database.DB

	if err := c.BodyParser(&data); err != nil {
		return err
	}
	var user models.User

	db.Where("name = ?", data["name"]).First(&user)

	if user.Id == 0 {
		return c.Status(400).JSON(models.Result{
			Status:  "error",
			Code:    500,
			Message: "Name not found",
			Data:    nil,
		})
	}

	if err := user.ComparePassword(data["password"]); err != nil {
		return c.Status(500).JSON(models.Result{
			Status:  "error",
			Code:    400,
			Message: "incorrect password",
			Data:    nil,
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

	return c.Status(200).JSON(models.Result{
		Status:  "success",
		Code:    200,
		Message: "User login successfully",
		Data:    user,
	})
}

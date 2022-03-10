package controllers

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"service-customer/database"
	"service-customer/models"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CreateUser(c *fiber.Ctx) error {
	type NewUser struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	db := database.DB
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(500).JSON(models.Result{
			Status:  "error",
			Code:    500,
			Message: "Input your user",
			Data:    nil,
		})
	}

	hash, err := hashPassword(user.Password)
	if err != nil {
		return c.Status(500).JSON(models.Result{
			Status:  "error",
			Code:    500,
			Message: "Couldn't hash password",
			Data:    nil,
		})
	}

	if user.Age <= 0 {
		return c.Status(500).JSON(models.Result{
			Status:  "error",
			Code:    500,
			Message: "Age is too small than 0",
			Data:    nil,
		})
	}

	user.Password = hash
	if err := db.Create(&user).Error; err != nil {
		return c.Status(500).JSON(models.Result{
			Status:  "error",
			Code:    500,
			Message: "Couldn't create user",
			Data:    nil,
		})
	}

	var i = 55
	if i-user.Age >= 30 {
		stockPercent := 72.5
		bondPercent := 21.5
		MMPercent := 100 - (stockPercent + bondPercent)

		riskProfile := models.RiskProfile{
			MmPercent:    float32(MMPercent),
			BondPercent:  float32(bondPercent),
			StockPercent: float32(stockPercent),
			UserId:       user.Id,
		}
		db.Create(&riskProfile)

	} else if i-user.Age >= 20 {
		stockPercent := 54.5
		bondPercent := 25.5
		MMPercent := 100 - (stockPercent + bondPercent)

		riskProfile := models.RiskProfile{
			MmPercent:    float32(MMPercent),
			BondPercent:  float32(bondPercent),
			StockPercent: float32(stockPercent),
			UserId:       user.Id,
		}
		db.Create(&riskProfile)

	} else if i-user.Age < 20 {
		stockPercent := 34.5
		bondPercent := 45.5
		MMPercent := 100 - (stockPercent + bondPercent)

		riskProfile := models.RiskProfile{
			MmPercent:    float32(MMPercent),
			BondPercent:  float32(bondPercent),
			StockPercent: float32(stockPercent),
			UserId:       user.Id,
		}
		db.Create(&riskProfile)
	}

	newUser := NewUser{
		Name: user.Name,
		Age:  user.Age,
	}

	return c.Status(200).JSON(models.Result{
		Status:  "success",
		Code:    200,
		Message: "User created successfully",
		Data:    newUser,
	})
}

package controllers

import (
	"github.com/gofiber/fiber/v2"
	"service-customer/database"
	"service-customer/models"
	"strconv"
)

func GetAllUser(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	take, _ := strconv.Atoi(c.Query("take", "10"))

	return c.JSON(models.Paginate(database.DB, &models.User{}, page, take))
}

func GetUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	db := database.DB

	risk := models.RiskProfile{
		Id: uint(id),
	}

	db.Find(&risk, id)
	//var d = risk.Id
	//if _, err := strconv.Atoi(strconv.Itoa(int(d))); err == nil {
	//	//if risk.Id != risk.d {
	//	return c.Status(404).JSON(models.Result{
	//		Status:  "error",
	//		Code:    404,
	//		Message: "No user found with ID",
	//		Data:    nil,
	//	})
	//}

	database.DB.Preload("User").Find(&risk)

	return c.Status(200).JSON(models.Result{
		Status:  "success",
		Code:    200,
		Message: "User found",
		Data:    risk,
	})
}

//func CreateUser(c *fiber.Ctx) error {
//	//id, _ := strconv.Atoi(c.Params("id"))
//	var data = make(map[int]int)
//	db := database.DB
//	//var user models.User
//	//var risk models.RiskProfile
//	if err := c.BodyParser(&data); err != nil {
//		return err
//	}
//
//	user := models.User{
//		Name: string(data["name"]),
//		Age:  data[age],
//		//RiskId: uint(id),
//	}
//	user.SetPassword(string(data["password"]))
//
//if err := c.BodyParser(&data); err != nil {
//	return c.Status(500).JSON(models.Result{
//		Status:  "error",
//		Code:    500,
//		Message: "Input your user",
//		Data:    nil,
//	})
//}
//
//if err := c.BodyParser(&data1); err != nil {
//	return c.Status(500).JSON(models.Result{
//		Status:  "error",
//		Code:    500,
//		Message: "Input your user 1",
//		Data:    nil,
//	})
//}
//type NewUserRequest struct {
//	SomeString int `validate:"regexp=\d"`
//}
//nur := NewUserRequest{SomeString: "JamesBond"}
//
//var name = user.Name
//var regex, _ = regexp.Compile(`\d`)
//var SomeString int `validate:"regexp=\d"`
//var isMatch = regex.FindAllString(name)
//if user.Name == NewUserRequest {
//	return c.Status(500).JSON(models.Result{
//		Status:  "error",
//		Code:    500,
//		Message: "Name cannot use numbers",
//		Data:    nil,
//	})
//}
//var use models.User
//age := use.Age
//user := models.User{
//	Name: data["Name"],
//	Age:  data1[age],
//}
//
//if err != nil {
//	return c.Status(500).JSON(models.Result{
//		Status:  "error",
//		Code:    500,
//		Message: "Couldn't hash password",
//		Data:    err,
//	})
//
//}
//
//if user.Age <= 0 {
//	return c.Status(500).JSON(models.Result{
//		Status:  "error",
//		Code:    500,
//		Message: "Age is too small than 0",
//		Data:    nil,
//	})
//}
//if err := db.Create(&user).Error; err != nil {
//
//	return c.Status(500).JSON(models.Result{
//		Status:  "error",
//		Code:    500,
//		Message: "Couldn't create user",
//		Data:    nil,
//	})
//}
//var i = 55
//if i-age >= 30 {
//	stockPercent := 72.5
//	bondPercent := 21.5
//	MMPercent := 100 - (stockPercent + bondPercent)
//
//	riskProfile := models.RiskProfile{
//		MmPercent:    float32(MMPercent),
//		BondPercent:  float32(bondPercent),
//		StockPercent: float32(stockPercent),
//		//UserId:       user.Id,
//	}
//	db.Create(&riskProfile)
//
//} else if i-age >= 20 {
//	stockPercent := 54.5
//	bondPercent := 25.5
//	MMPercent := 100 - (stockPercent + bondPercent)
//
//	riskProfile := models.RiskProfile{
//		MmPercent:    float32(MMPercent),
//		BondPercent:  float32(bondPercent),
//		StockPercent: float32(stockPercent),
//		//UserId:       user.Id,
//	}
//	db.Create(&riskProfile)
//} else if age < 20 {
//	stockPercent := 34.5
//	bondPercent := 45.5
//	MMPercent := 100 - (stockPercent + bondPercent)
//
//	riskProfile := models.RiskProfile{
//		MmPercent:    float32(MMPercent),
//		BondPercent:  float32(bondPercent),
//		StockPercent: float32(stockPercent),
//		//UserId:       user.Id,
//	}
//	db.Create(&riskProfile)
//}
//
//return c.Status(200).JSON(models.Result{
//	Status:  "success",
//	Code:    200,
//	Message: "User created successfully",
//	Data:    user,
//})
//}

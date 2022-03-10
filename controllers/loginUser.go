package controllers

//
//import (
//	"errors"
//	"gorm.io/gorm"
//
//	//"errors"
//	"github.com/gofiber/fiber/v2"
//	"github.com/golang-jwt/jwt"
//	"golang.org/x/crypto/bcrypt"
//	"service-customer/database"
//	"service-customer/models"
//
//	//"gorm.io/gorm"
//	"service-customer/config"
//	//"service-customer/models"
//	"time"
//)
//
//// CheckPasswordHash compare password with hash
//func CheckPasswordHash(password, hash string) bool {
//	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
//	return err == nil
//}
//
//func getUserByName(a string) (*models.User, error) {
//	db := database.DB
//	var user models.User
//	if err := db.Where(&models.User{Name: a}).Find(&user).Error; err != nil {
//		if errors.Is(err, gorm.ErrRecordNotFound) {
//			return nil, nil
//		}
//		return nil, err
//	}
//	return &user, nil
//}
//
//func getUserByPassword(i string) (*models.User, error) {
//	db := database.DB
//	var user models.User
//	if err := db.Where(&models.User{Password: i}).Find(&user).Error; err != nil {
//		if errors.Is(err, gorm.ErrRecordNotFound) {
//			return nil, nil
//		}
//		return nil, err
//	}
//	return &user, nil
//}
//
//func Login(c *fiber.Ctx) error {
//
//	type LoginInput struct {
//		Identity string `json:"identity"`
//		Name     string `json:"name"`
//		Password string `json:"password"`
//	}
//	type UserData struct {
//		Id       uint   `json:"id"`
//		Name     string `json:"name"`
//		Age      int    `json:"age"`
//		Password string `json:"password"`
//	}
//	var input LoginInput
//	var ud UserData
//
//	if err := c.BodyParser(&input); err != nil {
//		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on login request", "data": err})
//	}
//	identity := input.Identity
//	name1 := input.Name
//	pass1 := input.Password
//
//	age, err := getUserByAge(identity)
//	if err != nil {
//		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Error on name", "data": err})
//	}
//
//	pass, err := getUserByPassword(identity)
//	if err != nil {
//		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Error on password", "data": err})
//	}
//
//	if age == nil && pass == nil {
//		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "User not found", "data": err})
//	}
//
//	if age == nil {
//		ud = UserData{
//			Id:       pass.Id,
//			Name:     pass.Name,
//			Age:      pass.Age,
//			Password: pass.Password,
//		}
//	} else {
//		ud = UserData{
//			Id:       age.Id,
//			Name:     age.Name,
//			Age:      age.Age,
//			Password: age.Password,
//		}
//	}
//	if name1 != ud.Name {
//		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Error on password", "data": err})
//	}
//
//	if !CheckPasswordHash(pass1, ud.Password) {
//		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid password", "data": nil})
//	}
//
//	token := jwt.New(jwt.SigningMethodHS256)
//
//	claims := token.Claims.(jwt.MapClaims)
//	claims["name"] = ud.Name
//	claims["id"] = ud.Id
//	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
//
//	t, err := token.SignedString([]byte(config.Config("SECRET")))
//	if err != nil {
//		return c.SendStatus(fiber.StatusInternalServerError)
//	}
//
//	return c.JSON(fiber.Map{"status": "success", "message": "Success login", "data": t})
//}

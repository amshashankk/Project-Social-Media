package controllers

import (
	"github.com/amshashankk/GoAuth/dao"
	"github.com/amshashankk/GoAuth/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	passwd, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 10)

	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: passwd,
	}

	_, err := dao.UserData(&user)
	if err != nil {
		return c.JSON(err)
	} else {
		response := models.Response{
			Message: "User Registered SuccessFully",
			Success: true,
		}
		return c.JSON(response)

	}
}

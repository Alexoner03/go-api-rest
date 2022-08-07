package user

import (
	"github.com/gofiber/fiber/v2"
	"go-api-rest/modules/shared/utils"
)

func getUsers(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(GetUsers())
}

func createUser(c *fiber.Ctx) error {
	user := User{}
	err := c.BodyParser(&user)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	errValidate := utils.ValidateStruct(user)

	if errValidate != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errValidate)
	}

	userCreated, errCreated := CreateUser(user)

	if errCreated != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errCreated)
	}
	
	return c.Status(fiber.StatusOK).JSON(userCreated)
}

func updateUser(c *fiber.Ctx) error {
	return c.Status(200).JSON(User{})
}

func deleteUser(c *fiber.Ctx) error {
	return c.Status(200).JSON(User{})
}

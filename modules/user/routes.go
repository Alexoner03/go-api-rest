package user

import "github.com/gofiber/fiber/v2"

func loadRoutes(router fiber.Router) {
	router.Get("/", getUsers)
	router.Post("/", createUser)
	router.Put("/:id", updateUser)
	router.Delete("/:id", deleteUser)
}

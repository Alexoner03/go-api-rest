package types

import (
	"github.com/gofiber/fiber/v2"
)

type Module interface {
	SetRoutes(router fiber.Router)
}

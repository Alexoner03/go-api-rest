package user

import (
	"github.com/gofiber/fiber/v2"
)

type UserModule struct {
	Name string
}

func RegisterUserModule() *UserModule {
	return &UserModule{
		Name: "user",
	}
}

func (userModule *UserModule) SetRoutes(router fiber.Router) {
	group := router.Group("/" + userModule.Name)
	loadRoutes(group)
}
package modules

import (
	"github.com/gofiber/fiber/v2"
	"go-api-rest/modules/shared/types"
	"go-api-rest/modules/user"
)

func ConfigModules(app *fiber.App) {
	api := app.Group("/api")

	modules := []types.Module{
		user.RegisterUserModule(),
	}

	moduleLoader := types.NewModuleLoader(api, modules)
	moduleLoader.Load()
}

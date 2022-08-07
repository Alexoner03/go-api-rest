package types

import (
	"github.com/gofiber/fiber/v2"
)

type ModuleLoader struct {
	router  fiber.Router
	modules []Module
}

func NewModuleLoader(router fiber.Router, modules []Module) *ModuleLoader {
	return &ModuleLoader{
		router:  router,
		modules: modules,
	}
}

func (m *ModuleLoader) Load() {
	for _, module := range m.modules {
		module.SetRoutes(m.router)
	}
}

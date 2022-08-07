package main

import (
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"go-api-rest/modules"
	"go-api-rest/modules/database"
	"go-api-rest/modules/shared/utils"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	utils.CheckErrorAndDie(godotenv.Load())
	utils.CheckErrorAndDie(database.Connect())

	app := fiber.New()
	app.Use(recover.New())
	app.Use(logger.New())

	modules.ConfigModules(app)

	log.Fatal(app.Listen(":3000"))
}

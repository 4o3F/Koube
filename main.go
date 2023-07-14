package main

import (
	"github.com/4o3F/Koube/database"
	"github.com/4o3F/Koube/handlers"
	"github.com/4o3F/Koube/utils"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: "Koube",
	})

	utils.InitUtils()
	database.InitDatabase()
	handlers.InitHandlers()

	err := app.Listen(utils.KoubeConfig.Port)
	if err != nil {
		utils.KoubeLogger.Fatal().Err(err).Send()
	}
}

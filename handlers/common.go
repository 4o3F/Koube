package handlers

import (
	"github.com/4o3F/Koube/structs"
	"github.com/4o3F/Koube/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
	jsoniter "github.com/json-iterator/go"
)

func authSuccess(ctx *fiber.Ctx) error {
	return ctx.Next()
}

func authError(ctx *fiber.Ctx, err error) error {
	return sendCommonResponse(ctx, 403, "Wrong API key", nil)
}

func authValidator(ctx *fiber.Ctx, _ string) (bool, error) {
	apiAuthKey, ok := ctx.GetReqHeaders()["Authorization"]
	if ok {
		if apiAuthKey == utils.KoubeConfig.APIAuthKey {
			return true, nil
		} else {
			return false, nil
		}
	}
	return false, nil
}

func InitHandlers(app *fiber.App) {
	app.Use(keyauth.New(keyauth.Config{
		KeyLookup:      "header:Authorization",
		SuccessHandler: authSuccess,
		ErrorHandler:   authError,
		AuthScheme:     "",
		Validator:      authValidator,
	}))

	app.Post("/api/data/gen", genAudienceData)

	utils.KoubeLogger.Info().Msg("Init handlers complete")
}

func sendCommonResponse(ctx *fiber.Ctx, code int, message string, data map[string]interface{}) error {
	response := structs.KoubeResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
	json, err := jsoniter.Marshal(response)
	if err != nil {
		// THIS SHOULD NOT HAPPEN
		// If this happens, just stop the server and wait for further investigation
		utils.KoubeLogger.Fatal().Err(err).Send()
	}
	return ctx.Status(code).Send(json)
}

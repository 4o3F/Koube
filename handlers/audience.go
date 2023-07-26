package handlers

import (
	"github.com/4o3F/Koube/database"
	"github.com/4o3F/Koube/structs"
	"github.com/4o3F/Koube/utils"
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
	"time"
)

func checkEntranceCode(ctx *fiber.Ctx) error {
	var audience structs.KoubeAudience
	err := jsoniter.Unmarshal(ctx.Body(), &audience)
	if err != nil {
		return sendCommonResponse(ctx, 403, "非法输入", nil)
	}
	databaseAudience, err := database.GetAudienceByAid(audience.Aid)
	if err != nil {
		return sendCommonResponse(ctx, 500, "内部服务器错误", nil)
	}
	if databaseAudience.EntranceCode == audience.EntranceCode {
		return sendCommonResponse(ctx, 200, "匹配成功", nil)
	} else {
		return sendCommonResponse(ctx, 403, "匹配失败", nil)
	}
}

func checkVerifyCode(ctx *fiber.Ctx) error {
	var audience structs.KoubeAudience
	err := jsoniter.Unmarshal(ctx.Body(), &audience)
	if err != nil {
		return sendCommonResponse(ctx, 403, "非法输入", nil)
	}
	databaseAudience, err := database.GetAudienceByAid(audience.Aid)
	if err != nil {
		return sendCommonResponse(ctx, 500, "内部服务器错误", nil)
	}
	if databaseAudience.VerifyCode == audience.VerifyCode {
		return sendCommonResponse(ctx, 200, "匹配成功", nil)
	} else {
		return sendCommonResponse(ctx, 403, "匹配失败", nil)
	}
}

func getEntranceCodeConfig(ctx *fiber.Ctx) error {
	currentTime := time.Now().Unix()
	if utils.KoubeConfig.ShowStartTime-currentTime > utils.KoubeConfig.AccessAllowDuration {
		return sendCommonResponse(ctx, 403, "未到开放时间", nil)
	} else {
		return sendCommonResponse(ctx, 200, "获取成功", map[string]interface{}{
			"length": utils.KoubeConfig.MaxEntranceCodeLength,
			"salt":   utils.KoubeConfig.EntranceCodeSalt,
		})
	}
}

func audienceEnter(ctx *fiber.Ctx) error {
	var body struct {
		Aids []int `json:"aids"`
	}
	err := jsoniter.Unmarshal(ctx.Body(), &body)
	if err != nil {
		return sendCommonResponse(ctx, 403, "非法输入", nil)
	}
	for _, aid := range body.Aids {
		err = database.AudienceEnter(aid)
		if err != nil {
			return sendCommonResponse(ctx, 500, "内部服务器错误", nil)
		}
	}
	return sendCommonResponse(ctx, 200, "操作成功", nil)
}

package channelUser

import (
	"github.com/gofiber/fiber/v2"
	"youtube/apiType/commonType"
	"youtube/apiType/userChannelType"
	"youtube/models/channelUser"
	"youtube/pkg/response"
)

func JoinToChannel(ctx *fiber.Ctx) error {
	baseErrCode := "005"
	request := userChannelType.UserChannelRequest{}
	res := commonType.Response{}
	userID := ctx.Locals("UserID").(int64)
	errStr, code, err := request.DecodeValidate(ctx)
	if err != nil {
		return response.ErrorResponse(ctx, res, baseErrCode, "01", errStr, code)
	}
	exist, errStr, err := channelUser.Repo.ExistenceCheck(channelUser.ChannelUser{
		UserID:    userID,
		ChannelID: request.ChannelID,
	})
	if err != nil {
		return response.ErrorResponse(ctx, res, baseErrCode, "02", errStr, 500)
	}
	if exist {
		return response.ErrorResponse(ctx, res, baseErrCode, "03", errStr, 200)
	}
	errStr, err = channelUser.Repo.Insert(channelUser.ChannelUser{
		UserID:    userID,
		ChannelID: request.ChannelID,
	})
	if err != nil {
		return response.ErrorResponse(ctx, res, baseErrCode, "04", errStr, 500)
	}
	return response.SuccessResponse(ctx, res)
}

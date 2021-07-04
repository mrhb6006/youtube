package channel

import (
	"github.com/gofiber/fiber/v2"
	"youtube/apiType/channelType"
	"youtube/apiType/commonType"
	channelRepo "youtube/models/chennel"
	"youtube/pkg/response"
)

func DeleteChannel(ctx *fiber.Ctx) error {
	baseErrCode := "020"
	request := channelType.DeleteChannelRequest{}
	res := commonType.Response{}
	errStr, code, err := request.DecodeValidate(ctx)
	if err != nil {
		return response.ErrorResponse(ctx, res, baseErrCode, "01", errStr, code)
	}
	errStr, err = channelRepo.Repo.DeleteChannel(request.ChannelID)
	if err != nil {
		return response.ErrorResponse(ctx, res, baseErrCode, "02", errStr, 500)
	}
	return response.SuccessResponse(ctx, res)
}

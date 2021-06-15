package channelVideo

import (
	"github.com/gofiber/fiber/v2"
	"youtube/apiType/channelVideoType"
	"youtube/apiType/commonType"
	"youtube/models/channelVideo"
	"youtube/pkg/response"
)

func DeleteVideoFromChannel(ctx *fiber.Ctx) error {
	baseErrCode := "004"
	request := channelVideoType.ChannelVideoRequest{}
	res := commonType.Response{}
	errStr, code, err := request.DecodeValidate(ctx)
	if err != nil {
		return response.ErrorResponse(ctx, res, baseErrCode, "01", errStr, code)
	}
	exist, errStr, err := channelVideo.Repo.ExistenceCheck(channelVideo.ChannelVideo{
		ChannelID: request.ChannelID,
		VideoID:   request.VideoID,
	})
	if err != nil {
		return response.ErrorResponse(ctx, res, baseErrCode, "02", errStr, 500)
	}
	if !exist {
		return response.ErrorResponse(ctx, res, baseErrCode, "03", errStr, 200)
	}
	errStr, err = channelVideo.Repo.Delete(channelVideo.ChannelVideo{
		ChannelID: request.ChannelID,
		VideoID:   request.VideoID,
	})
	if err != nil {
		return response.ErrorResponse(ctx, res, baseErrCode, "04", errStr, 500)
	}
	return response.SuccessResponse(ctx, res)
}

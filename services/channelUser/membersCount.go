package channelUser

import (
	"github.com/gofiber/fiber/v2"
	"youtube/apiType/commonType"
	"youtube/apiType/userChannelType"
	"youtube/models/channelUser"
	"youtube/pkg/response"
)

func ChannelMembersCount(ctx *fiber.Ctx) error {
	baseErrCode := "013"
	request := userChannelType.UserChannelRequest{}
	commentLikeCountResponse := userChannelType.CommentLikeCountResponse{}
	res := commonType.Response{}
	errStr, code, err := request.DecodeValidate(ctx)
	if err != nil {
		return response.ErrorResponse(ctx, res, baseErrCode, "01", errStr, code)
	}
	count, errStr, err := channelUser.Repo.GetMembersCount(request.ChannelID)
	if err != nil {
		return response.ErrorResponse(ctx, res, baseErrCode, "02", errStr, 500)
	}
	commentLikeCountResponse.MemberCount = count
	res.Res = commentLikeCountResponse
	return response.SuccessResponse(ctx, res)
}

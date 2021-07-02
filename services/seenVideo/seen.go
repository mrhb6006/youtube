package seenVideo

import (
	"github.com/gofiber/fiber/v2"
	"youtube/apiType/commonType"
	"youtube/apiType/seenVideoType"
	"youtube/models/seenVideo"
	"youtube/pkg/response"
)

func SeenVideo(ctx *fiber.Ctx) error {
	baseErrCode := "007"
	request := seenVideoType.SeenVideoRequest{}
	res := commonType.Response{}
	userID := ctx.Locals("UserID").(int64)
	errStr, code, err := request.DecodeValidate(ctx)
	if err != nil {
		return response.ErrorResponse(ctx, res, baseErrCode, "01", errStr, code)
	}

	errStr, err = seenVideo.Repo.Insert(seenVideo.Seen{
		UserID:  userID,
		VideoID: request.VideoID,
	})
	if err != nil {
		return response.ErrorResponse(ctx, res, baseErrCode, "02", errStr, 500)
	}
	return response.SuccessResponse(ctx, res)
}

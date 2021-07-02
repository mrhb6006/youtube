package seenVideo

import (
	"github.com/gofiber/fiber/v2"
	"youtube/apiType/commonType"
	"youtube/apiType/seenVideoType"
	"youtube/models/seenVideo"
	"youtube/pkg/response"
)

func CountVideoSeen(ctx *fiber.Ctx) error {
	baseErrCode := "008"
	request := seenVideoType.SeenVideoRequest{}
	countVideoSeensResponse := seenVideoType.CountVideoSeensResponse{}
	res := commonType.Response{}
	errStr, code, err := request.DecodeValidate(ctx)
	if err != nil {
		return response.ErrorResponse(ctx, res, baseErrCode, "01", errStr, code)
	}
	count, errStr, err := seenVideo.Repo.CountVideoSeens(request.VideoID)
	if err != nil {
		return response.ErrorResponse(ctx, res, baseErrCode, "02", errStr, 500)
	}
	countVideoSeensResponse.Count = count
	res.Res = countVideoSeensResponse
	return response.SuccessResponse(ctx, res)
}

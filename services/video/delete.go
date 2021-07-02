package video

import (
	"github.com/gofiber/fiber/v2"
	"youtube/apiType/commonType"
	"youtube/apiType/videoType"
	"youtube/models/video"
	"youtube/pkg/response"
)

func DeleteVideo(ctx *fiber.Ctx) error {
	baseErrCode := "014"
	request := videoType.DeleteVideoRequest{}
	res := commonType.Response{}
	errStr, code, err := request.DecodeValidate(ctx)
	if err != nil {
		return response.ErrorResponse(ctx, res, baseErrCode, "01", errStr, code)
	}
	errStr, err = video.Repo.Delete(request.VideoID)
	if err != nil {
		return response.ErrorResponse(ctx, res, baseErrCode, "02", errStr, 500)
	}
	return response.SuccessResponse(ctx, res)
}

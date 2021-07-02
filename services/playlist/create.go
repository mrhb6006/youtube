package playlist

import (
	"github.com/gofiber/fiber/v2"
	"youtube/apiType/commentType"
	"youtube/apiType/commonType"
	"youtube/apiType/playListType"
	"youtube/models/playlist"
	"youtube/pkg/response"
)

func Create(ctx *fiber.Ctx) error {
	baseErrCode := "072"
	request := playListType.CreatePlayListRequest{}
	res := commonType.Response{}
	errStr, code, err := request.DecodeValidate(ctx)
	if err != nil {
		return response.ErrorResponse(ctx, res, baseErrCode, "01", errStr, code)
	}
}

package playlist

import (
	"github.com/gofiber/fiber/v2"
	"youtube/apiType/commonType"
	"youtube/apiType/playListType"
	"youtube/models/playlist"
	"youtube/pkg/response"
)

func DeletePlayList(ctx *fiber.Ctx) error {
	baseErrCode := "076"
	request := playListType.DeletePlayListRequest{}
	res := commonType.Response{}
	userID := ctx.Locals("UserID").(int64)
	errStr, code, err := request.DecodeValidate(ctx)
	if err != nil {
		return response.ErrorResponse(ctx, res, baseErrCode, "01", errStr, code)
	}
	list, exist, errStr, err := playlist.Repo.GetByID(request.PlayListID)
	if err != nil {
		return response.ErrorResponse(ctx, res, baseErrCode, "02", errStr, 500)
	}
	if !exist {
		return response.ErrorResponse(ctx, res, baseErrCode, "03", "01", 200)
	}
	if list.CreatorID != userID {
		return response.ErrorResponse(ctx, res, baseErrCode, "04", "01", 200)
	}
	if list.IsDefault {
		return response.ErrorResponse(ctx, res, baseErrCode, "05", "01", 200)
	}

	errStr, err = playlist.Repo.Delete(request.PlayListID)
	if err != nil {
		return response.ErrorResponse(ctx, res, baseErrCode, "06", errStr, 500)
	}
	return response.SuccessResponse(ctx, res)
}

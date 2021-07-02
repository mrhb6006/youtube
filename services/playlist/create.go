package playlist

import (
	"github.com/gofiber/fiber/v2"
	"youtube/apiType/commonType"
	"youtube/apiType/playListType"
	"youtube/models/playlist"
	"youtube/pkg/response"
)

func CreatePlayList(ctx *fiber.Ctx) error {
	baseErrCode := "072"
	request := playListType.CreatePlayListRequest{}
	createPlayListResponse := playListType.CreatePlayListResponse{}
	res := commonType.Response{}
	userID := ctx.Locals("UserID").(int64)
	errStr, code, err := request.DecodeValidate(ctx)
	if err != nil {
		return response.ErrorResponse(ctx, res, baseErrCode, "01", errStr, code)
	}
	_, exist, errStr, err := playlist.Repo.GetByName(request.Name, userID)
	if err != nil {
		return response.ErrorResponse(ctx, res, baseErrCode, "02", errStr, 500)
	}
	if exist {
		return response.ErrorResponse(ctx, res, baseErrCode, "03", "01", 200)
	}
	insertedID, errStr, err := playlist.Repo.Insert(playlist.PlayList{
		Name:      request.Name,
		IsPublic:  request.IsPublic,
		CreatorID: userID,
	})
	if err != nil {
		return response.ErrorResponse(ctx, res, baseErrCode, "02", errStr, 500)
	}
	createPlayListResponse.PlayListID = insertedID
	res.Res = createPlayListResponse
	return response.SuccessResponse(ctx, res)
}

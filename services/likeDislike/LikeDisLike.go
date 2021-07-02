package likeDislike

import (
	"github.com/gofiber/fiber/v2"
	"youtube/apiType/commonType"
	"youtube/apiType/likeVideoType"
	"youtube/models/likeDislike"
	"youtube/pkg/response"
)

func LikeDislikeVideo(ctx *fiber.Ctx) error {
	baseErrCode := "009"
	request := likeVideoType.LikeVideoRequest{}
	res := commonType.Response{}
	userID := ctx.Locals("UserID").(int64)
	errStr, code, err := request.DecodeValidate(ctx)
	if err != nil {
		return response.ErrorResponse(ctx, res, baseErrCode, "01", errStr, code)
	}
	exist, errStr, err := likeDislike.Repo.ChechExist(userID, request.VideoID)
	if err != nil {
		return response.ErrorResponse(ctx, res, baseErrCode, "03", errStr, 500)
	}

	if exist {
		errStr, err = likeDislike.Repo.UpdateAction(likeDislike.Like{
			VideoID: request.VideoID,
			UserID:  userID,
			Action:  request.Action,
		})
		if err != nil {
			return response.ErrorResponse(ctx, res, baseErrCode, "04", errStr, 500)
		}
	} else {
		errStr, err = likeDislike.Repo.Insert(likeDislike.Like{
			VideoID: request.VideoID,
			UserID:  userID,
			Action:  request.Action,
		})
		if err != nil {
			return response.ErrorResponse(ctx, res, baseErrCode, "05", errStr, 500)
		}
	}
	return response.SuccessResponse(ctx, res)
}

package likeDislike

import (
	"github.com/gofiber/fiber/v2"
	"youtube/apiType/commonType"
	"youtube/apiType/likeVideoType"
	"youtube/models/likeDislikeVideo"
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
	exist, errStr, err := likeDislikeVideo.Repo.ChechExist(userID, request.VideoID)
	if err != nil {
		return response.ErrorResponse(ctx, res, baseErrCode, "03", errStr, 500)
	}

	if exist {
		errStr, err = likeDislikeVideo.Repo.UpdateAction(likeDislikeVideo.LikeVideo{
			VideoID: request.VideoID,
			UserID:  userID,
			Action:  request.Action,
		})
		if err != nil {
			return response.ErrorResponse(ctx, res, baseErrCode, "04", errStr, 500)
		}
	} else {
		errStr, err = likeDislikeVideo.Repo.Insert(likeDislikeVideo.LikeVideo{
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

func VideoLikesCount(ctx *fiber.Ctx) error {
	baseErrCode := "011"
	request := likeVideoType.LikeVideoCountRequest{}
	likeCountResponse := likeVideoType.LikeVideoCountResponse{}
	res := commonType.Response{}
	errStr, code, err := request.DecodeValidate(ctx)
	if err != nil {
		return response.ErrorResponse(ctx, res, baseErrCode, "01", errStr, code)
	}
	count, errStr, err := likeDislikeVideo.Repo.GetVideoLikesCount(request.Action, request.VideoID)
	if err != nil {
		return response.ErrorResponse(ctx, res, baseErrCode, "02", errStr, 500)
	}
	likeCountResponse.LikeCount = count
	res.Res = likeCountResponse
	return response.SuccessResponse(ctx, res)
}

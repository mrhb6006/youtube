package likeDislike

import (
	"github.com/gofiber/fiber/v2"
	"youtube/apiType/commonType"
	"youtube/apiType/likeCommentType"
	"youtube/models/likeDislikeComment"
	"youtube/pkg/response"
)

func LikeDislikeComment(ctx *fiber.Ctx) error {
	baseErrCode := "010"
	request := likeCommentType.LikeCommentRequest{}
	res := commonType.Response{}
	userID := ctx.Locals("UserID").(int64)
	errStr, code, err := request.DecodeValidate(ctx)
	if err != nil {
		return response.ErrorResponse(ctx, res, baseErrCode, "01", errStr, code)
	}
	exist, errStr, err := likeDislikeComment.Repo.ChechExist(userID, request.CommentID)
	if err != nil {
		return response.ErrorResponse(ctx, res, baseErrCode, "02", errStr, 500)
	}

	if exist {
		errStr, err = likeDislikeComment.Repo.UpdateAction(likeDislikeComment.Like{
			CommentID: request.CommentID,
			UserID:    userID,
			Action:    request.Action,
		})
		if err != nil {
			return response.ErrorResponse(ctx, res, baseErrCode, "03", errStr, 500)
		}
	} else {
		errStr, err = likeDislikeComment.Repo.Insert(likeDislikeComment.Like{
			CommentID: request.CommentID,
			UserID:    userID,
			Action:    request.Action,
		})
		if err != nil {
			return response.ErrorResponse(ctx, res, baseErrCode, "04", errStr, 500)
		}
	}
	return response.SuccessResponse(ctx, res)
}

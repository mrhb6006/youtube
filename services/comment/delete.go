package comment

import (
	"github.com/gofiber/fiber/v2"
	"youtube/apiType/commentType"
	"youtube/apiType/commonType"
	commentRepo "youtube/models/comment"
	"youtube/pkg/response"
)

func DeleteComment(ctx *fiber.Ctx) error {
	baseErrCode := "081"
	request := commentType.DeleteCommentRequest{}
	res := commonType.Response{}
	errStr, code, err := request.DecodeValidate(ctx)
	if err != nil {
		return response.ErrorResponse(ctx, res, baseErrCode, "01", errStr, code)
	}
	deletedID, errStr, err := commentRepo.Repo.Delete(request.CommentID)
	if err != nil {
		return response.ErrorResponse(ctx, res, baseErrCode, "02", errStr, 500)
	}
	if deletedID == 0 {
		return response.ErrorResponse(ctx, res, baseErrCode, "03", "03", 200)
	}
	return response.SuccessResponse(ctx, res)
}

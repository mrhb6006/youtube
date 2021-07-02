package comment

import (
	"github.com/gofiber/fiber/v2"
	"strings"
	"youtube/apiType/commentType"
	"youtube/apiType/commonType"
	commentRepo "youtube/models/comment"
	"youtube/pkg/response"
)

func Write(ctx *fiber.Ctx) error {
	baseErrCode := "080"
	request := commentType.WriteCommentRequest{}
	commentResponse := commentType.WriteCommentResponse{}
	res := commonType.Response{}
	userID := ctx.Locals("UserID").(int64)
	errStr, code, err := request.DecodeValidate(ctx)
	if err != nil {
		return response.ErrorResponse(ctx, res, baseErrCode, "01", errStr, code)
	}
	comment := commentRepo.Comment{
		Text:    strings.TrimSpace(request.Text),
		ReplyID: request.ReplyID,
		VideoID: request.VideoID,
		UserID:  userID,
	}
	if request.ReplyID != 0 { // reply a comment
		comment.ReplyID = request.ReplyID
	}
	id, errStr, err := commentRepo.Repo.Insert(comment)
	if err != nil {
		return response.ErrorResponse(ctx, res, baseErrCode, "02", errStr, 500)
	}
	commentResponse.CommentID = id
	res.Res = commentResponse
	return response.SuccessResponse(ctx, res)
}

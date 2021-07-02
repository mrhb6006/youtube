package comment

import (
	"github.com/gofiber/fiber/v2"
	"youtube/apiType/commentType"
	"youtube/apiType/commonType"
	"youtube/apiType/videoType"
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
	if request.ReplyID == 0 {

	} else {

	}
}

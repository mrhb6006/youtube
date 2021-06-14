package response

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"youtube/apiType/commonType"
	"youtube/pkg/errorHandler"
)

func ErrorResponse(ctx *fiber.Ctx, res commonType.Response, baseErrorCode, SectionNo, errStr string, code int) error {
	errCode := baseErrorCode + SectionNo + errStr
	res.Err = errorHandler.CreateError(ctx, errCode, code)
	if code > 200 {
		return res.Err
	}
	return SuccessResponse(ctx, res)

}
func SuccessResponse(ctx *fiber.Ctx, res commonType.Response) error {
	res.Status = "success"
	if len(res.Err["code"]) > 0 {
		res.Status = "error"
	}
	data, err := json.Marshal(res)
	if err != nil {
		zap.L().Error("marshal_err", zap.Int("status", 500))
		return errorHandler.CreateError(ctx, "9999999", 500)
	}
	_, _ = ctx.Write(data)
	return nil
}

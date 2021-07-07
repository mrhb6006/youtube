package errorHandler

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"youtube/apiType/commonType"
)

func DefaultErrHandler(ctx *fiber.Ctx, err error) error {
	code := "0000000"
	if e, ok := err.(commonType.Error); ok {
		code = e["code"]
	}
	ctx.Set(fiber.HeaderContentType, fiber.MIMEOctetStream)
	comm := commonType.Response{
		Status: "error",
		Err: commonType.Error{
			"code":    code,
			"message": err.Error(),
		},
	}
	b, er := json.Marshal(comm)
	if er != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(GetMessage(code, ctx.Get("language", "en")))
	}
	return ctx.Send(b)
}

func CreateError(ctx *fiber.Ctx, err string, code int) commonType.Error {
	var e = make(commonType.Error)
	e["code"] = err
	ctx.Status(code)
	if code > 200 {
		err = strconv.Itoa(code)
	}
	e["message"] = GetMessage(err, ctx.Get("language", "en"))
	return e
}

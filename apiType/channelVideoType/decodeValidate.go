package channelVideoType

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"youtube/pkg/validate"
)

func (r *ChannelVideoRequest) DecodeValidate(ctx *fiber.Ctx) (errStr string, responseCode int, err error) {
	err = json.Unmarshal(ctx.Body(), r)
	if err != nil {
		return "01", 400, err
	}
	err = validate.Struct(r)
	if err != nil {
		customError := err.(validator.ValidationErrors)
		switch customError[0].StructField() + "," + customError[0].ActualTag() {
		case "ChannelID,required":
			return "02", 400, err
		case "ChannelID,gt":
			return "03", 400, err
		case "VideoID,required":
			return "04", 400, err
		case "VideoID,gt":
			return "05", 400, err
		}
	}
	return "", 200, nil
}

package channelType

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"youtube/pkg/validate"
)

func (r *CreateChannelRequest) DecodeValidate(ctx *fiber.Ctx) (string, int, error) {
	err := json.Unmarshal(ctx.Body(), r)
	if err != nil {
		return "01", 400, err
	}
	err = validate.Struct(r)
	if err != nil {
		customError := err.(validator.ValidationErrors)
		switch customError[0].StructField() + "," + customError[0].ActualTag() {
		case "Name,min":
			return "02", 400, err
		case "Name,required":
			return "03", 400, err
		case "Description,required":
			return "04", 400, err
		case "Description,max":
			return "05", 400, err
		case "Avatar,required":
			return "06", 400, err
		}
	}

	return "", 200, nil
}

func (r *DeleteChannelRequest) DecodeValidate(ctx *fiber.Ctx) (errStr string, responseCode int, err error) {
	err = json.Unmarshal(ctx.Body(), r)
	if err != nil {
		return "01", 400, err
	}
	err = validate.Struct(r)
	if err != nil {
		customError := err.(validator.ValidationErrors)
		switch customError[0].StructField() + "," + customError[0].ActualTag() {
		case "ChannelID,gt":
			return "02", 400, err
		case "ChannelID,required":
			return "03", 400, err
		}
	}
	return "", 200, nil
}

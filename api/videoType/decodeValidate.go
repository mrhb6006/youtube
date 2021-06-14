package videoType

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"youtube/pkg/validate"
)

func (r *UploadVideoRequest) DecodeValidate(ctx *fiber.Ctx) (string, int, error) {
	err := json.Unmarshal(ctx.Body(), r)
	if err != nil {
		return "01", 400, err
	}
	err = validate.Struct(r)
	if err != nil {
		customError := err.(validator.ValidationErrors)
		switch customError[0].StructField() + "," + customError[0].ActualTag() {
		case "ChannelID,required":
			return "02", 400, err
		case "Title,required":
			return "03", 400, err
		case "Duration,required":
			return "04", 400, err
		case "Thumbnail,required":
			return "05", 400, err
		case "Video,required":
			return "06", 400, err
		case "ChannelID,gt":
			return "07", 400, err
		case "Description,max":
			return "08", 400, err
		case "Duration,gt":
			return "09", 400, err
		}
	}
	return "", 200, nil
}

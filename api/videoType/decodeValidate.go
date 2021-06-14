package videoType

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"youtube/pkg/validate"
)

func (r *UploadVideoRequest) DecodeValidate(ctx *fiber.Ctx) (int, string, error) {
	err := json.Unmarshal(ctx.Body(), r)
	if err != nil {
		return 400, "01", err
	}
	err = validate.Struct(r)
	if err != nil {
		customError := err.(validator.ValidationErrors)
		switch customError[0].StructField() + "," + customError[0].ActualTag() {
		case "ChannelID,required":
			return 400, "02", err
		case "Title,required":
			return 400, "03", err
		case "Duration,required":
			return 400, "04", err
		case "Thumbnail,required":
			return 400, "05", err
		case "VideoBase64,required":
			return 400, "06", err
		case "ChannelID,gt":
			return 400, "07", err
		case "Description,max":
			return 400, "08", err
		case "Duration,gt":
			return 400, "09", err
		}
	}
	return 200, "", nil
}

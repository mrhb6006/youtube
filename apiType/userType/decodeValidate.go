package userType

import (
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"net/mail"
	"strings"
	"youtube/pkg/validate"
)

const MinNameLength = 5
const MinPassLength = 8

func (user *RegisterRequest) DecodeValidate(ctx *fiber.Ctx) (string, int, error) {
	err := json.Unmarshal(ctx.Body(), user)
	if err != nil {
		return "01", 400, err
	}
	err = validate.Struct(user)
	if err != nil {
		customError := err.(validator.ValidationErrors)
		switch customError[0].StructField() + "," + customError[0].ActualTag() {
		case "Username,required":
			return "02", 400, err
		case "Password,required":
			return "03", 400, err
		}
	}

	user.Username = strings.TrimSpace(user.Username)
	user.Password = strings.TrimSpace(user.Password)
	user.Email = strings.TrimSpace(user.Email)

	if len(user.Username) < MinNameLength {
		return "04", 200, errors.New("invalid user param")
	} else if len(user.Password) < MinPassLength {
		return "05", 200, errors.New("invalid user param")
	}
	if user.Email != "" {
		if len(user.Email) < 5 {
			return "06", 200, errors.New("invalid user param")
		}
		_, err := mail.ParseAddress(user.Email)
		if err != nil {
			return "07", 200, errors.New("invalid user param")
		}
	}
	return "", 200, nil
}

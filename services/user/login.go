package user

import (
	"crypto/md5"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"youtube/apiType/commonType"
	"youtube/apiType/userType"
	"youtube/models/user"
	"youtube/pkg/auth"
	"youtube/pkg/response"
)

func Login(ctx *fiber.Ctx) error {
	baseErrCode := "099"
	request := userType.LoginRequest{}
	registerResponse := userType.RegisterResponse{}
	res := commonType.Response{}

	errStr, code, err := request.DecodeValidate(ctx)
	if err != nil {
		return response.ErrorResponse(ctx, res, baseErrCode, "01", errStr, code)
	}

	foundedUser, found, errStr, err := user.Repo.GetByUserName(request.Username)
	if err != nil {
		return response.ErrorResponse(ctx, res, baseErrCode, "02", errStr, 500)
	}

	if !found {
		return response.ErrorResponse(ctx, res, baseErrCode, "03", "01", 200)
	}

	byteArray := md5.Sum([]byte(request.Password))
	hashPassStr := fmt.Sprintf("%x", byteArray)

	if foundedUser.Password != hashPassStr {
		return response.ErrorResponse(ctx, res, baseErrCode, "04", "01", 200)
	}

	registerResponse.Token, err = auth.CreateToken(foundedUser.ID)
	if err != nil {
		return response.ErrorResponse(ctx, res, baseErrCode, "05", "01", 500)
	}
	res.Res = registerResponse
	return response.SuccessResponse(ctx, res)
}

package channel

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"strconv"
	"time"
	"youtube/apiType/channelType"
	"youtube/apiType/commonType"
	channelRepo "youtube/models/chennel"
	"youtube/pkg/response"
	"youtube/pkg/storageHandler"
)

func Create(ctx *fiber.Ctx) error {
	baseErrCode := "002"
	request := channelType.CreateChannelRequest{}
	createChannelResponse := channelType.CreateChannelResponse{}
	res := commonType.Response{}
	userID := ctx.Locals("UserID").(int64)
	errStr, code, err := request.DecodeValidate(ctx)
	if err != nil {
		return response.ErrorResponse(ctx, res, baseErrCode, "01", errStr, code)
	}
	avatarPath, err := storageHandler.SaveImage(request.Avatar, "a"+strconv.FormatInt(time.Now().UnixNano(), 10))
	if err != nil {
		zap.L().Error("ERROR:", zap.Any("err:", err.Error()))
		return response.ErrorResponse(ctx, res, baseErrCode, "02", "01", 500)
	}
	_, exist, errStr, err := channelRepo.Repo.GetByName(request.Name)
	if err != nil {
		return response.ErrorResponse(ctx, res, baseErrCode, "03", errStr, 500)
	}
	if exist {
		return response.ErrorResponse(ctx, res, baseErrCode, "04", "01", 200)
	}

	id, errStr, err := channelRepo.Repo.CreateChannel(channelRepo.Channel{
		Name:        request.Name,
		Description: request.Description,
		Avatar:      avatarPath,
		CreatorID:   userID,
	})

	if err != nil {
		return response.ErrorResponse(ctx, res, baseErrCode, "05", "01", 500)
	}
	createChannelResponse.Id = id
	res.Res = createChannelResponse
	return response.SuccessResponse(ctx, res)
}

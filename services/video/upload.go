package video

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"strconv"
	"time"
	"youtube/apiType/commonType"
	"youtube/apiType/videoType"
	channelVideoRepo "youtube/models/channelVideo"
	channelRepo "youtube/models/chennel"
	"youtube/models/storage"
	"youtube/models/video"
	"youtube/pkg/response"
	"youtube/pkg/storageHandler"
)

func Upload(ctx *fiber.Ctx) error {
	baseErrCode := "001"
	request := videoType.UploadVideoRequest{}
	uploadVideoResponse := videoType.UploadVideoResponse{}
	res := commonType.Response{}
	errStr, code, err := request.DecodeValidate(ctx)
	if err != nil {
		return response.ErrorResponse(ctx, res, baseErrCode, "01", errStr, code)
	}

	channel, found, errStr, err := channelRepo.Repo.GetByID(request.ChannelID)
	if err != nil {
		return response.ErrorResponse(ctx, res, baseErrCode, "02", errStr, 500)
	}
	if !found {
		return response.ErrorResponse(ctx, res, baseErrCode, "03", "01", 200)
	}

	//todo check user access to upload

	thumbnailPath, err := storageHandler.SaveImage(request.Thumbnail, "t"+strconv.FormatInt(time.Now().UnixNano(), 10))
	if err != nil {
		zap.L().Error("save image error", zap.Any("err:", err))
		return response.ErrorResponse(ctx, res, baseErrCode, "04", "01", 500)
	}
	VideoPath, err := storageHandler.SaveVideo(request.Video, "v"+strconv.FormatInt(time.Now().UnixNano(), 10))
	if err != nil {
		zap.L().Error("save video error", zap.Any("err:", err))
		return response.ErrorResponse(ctx, res, baseErrCode, "05", "01", 500)
	}

	storageID, errStr, err := storage.Repo.Insert(storage.Storage{
		Path: VideoPath,
	})
	if err != nil {
		return response.ErrorResponse(ctx, res, baseErrCode, "06", "01", 500)
	}

	videoId, errStr, err := video.Repo.Insert(video.Video{
		Title:       request.Title,
		Description: request.Description,
		Duration:    request.Duration,
		Thumbnail:   thumbnailPath,
		StorageID:   storageID,
	})
	if err != nil {
		return response.ErrorResponse(ctx, res, baseErrCode, "07", "01", 500)
	}
	errStr, err = channelVideoRepo.Repo.Insert(channelVideoRepo.ChannelVideo{
		ChannelID: channel.ID,
		VideoID:   videoId,
	})
	if err != nil {
		return response.ErrorResponse(ctx, res, baseErrCode, "08", "01", 500)
	}
	uploadVideoResponse.ID = videoId
	res.Res = uploadVideoResponse
	return response.SuccessResponse(ctx, res)
}

package channelVideoType

type ChannelVideoRequest struct {
	ChannelID int64 `json:"channel_id" validate:"required,gt=0"`
	VideoID   int64 `json:"video_id" validate:"required,gt=0"`
}

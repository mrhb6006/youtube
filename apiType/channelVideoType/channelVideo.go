package channelVideoType

type ChannelVideoRequest struct {
	ChannelID int64 `json:"channel_id" validate:"required,gt=0"`
	VideoID   int64 `json:"video_id" validate:"required,gt=0"`
}

type ChannelVideoResponse struct {
	VideoAdded bool `json:"video_added"`
}

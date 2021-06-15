package videoChannelType

type VideoChannelRequest struct {
	ChannelID int64 `json:"channel_id" validate:"required,gt=0"`
	VideoID   int64 `json:"video_id" validate:"required,gt=0"`
}

type VideoChannelResponse struct {
	VideoAdded bool `json:"video_added"`
}

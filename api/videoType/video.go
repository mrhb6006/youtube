package videoType

type UploadVideoRequest struct {
	ChannelID   int64  `json:"channel_id" validate:"required,gt=0"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"max=256"`
	Duration    int64  `json:"duration" validate:"required,gt=0"`
	Thumbnail   string `json:"thumbnail" validate:"required"`
	VideoBase64 string `json:"video_base64" validate:"required"`
}

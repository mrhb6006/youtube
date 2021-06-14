package videoType

type UploadVideoRequest struct {
	ChannelID   int64  `json:"channel_id" validate:"required,gt=0"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"max=256"`
	Duration    int64  `json:"duration" validate:"required,gt=0"`
	Thumbnail   string `json:"thumbnail" validate:"required"`
	Video       string `json:"video" validate:"required"`
}

type UploadVideoResponse struct {
	ID int64 `json:"id"`
}

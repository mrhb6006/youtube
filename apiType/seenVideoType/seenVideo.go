package seenVideoType

type SeenVideoRequest struct {
	VideoID int64 `json:"video_id" validate:"required,gt=0"`
}
type CountVideoSeensResponse struct {
	Count int64 `json:"video_seens_count"`
}

package likeVideoType

type LikeVideoRequest struct {
	Action  int64 `json:"action" validate:"oneof=0 1"`
	VideoID int64 `json:"video_id" validate:"gt=0,required"`
}

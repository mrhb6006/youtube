package likeVideoType

type LikeVideoRequest struct {
	Action  int64 `json:"action" validate:"oneof=0 1"`
	VideoID int64 `json:"video_id" validate:"gt=0,required"`
}

type LikeVideoCountRequest struct {
	Action  int64 `json:"action" validate:"oneof=0 1"`
	VideoID int64 `json:"video_id" validate:"gt=0,required"`
}

type LikeVideoCountResponse struct {
	LikeCount int64 `json:"like_count"`
}

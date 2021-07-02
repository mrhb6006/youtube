package likeCommentType

type LikeCommentRequest struct {
	Action    int64 `json:"action" validate:"oneof=0 1"`
	CommentID int64 `json:"comment_id" validate:"gt=0,required"`
}

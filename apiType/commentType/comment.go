package commentType

type WriteCommentRequest struct {
	Text    string `json:"text" validate:"required"`
	ReplyID int64  `json:"reply_id" validate:"gt=0"`
	VideoID int64  `json:"video_id" validate:"gt=0,required"`
}

type WriteCommentResponse struct {
	CommentID int64 `json:"comment_id"`
}

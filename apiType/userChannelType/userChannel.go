package userChannelType

// UserChannelRequest : request format for join and left channels
type UserChannelRequest struct {
	ChannelID int64 `json:"channel_id" validate:"required,gt=0"`
}

type CommentLikeCountResponse struct {
	MemberCount int64 `json:"member_count"`
}

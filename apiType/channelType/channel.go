package channelType

type CreateChannelRequest struct {
	Name        string `json:"name" validate:"min=3,required"`
	Description string `json:"description" validate:"max=256,required"`
	Avatar      string `json:"avatar" validate:"required"`
}

type CreateChannelResponse struct {
	Id int64 `json:"id"`
}

package playListType

type CreatePlayListRequest struct {
	Name     string `json:"name" validate:"required"`
	IsPublic bool   `json:"is_public"`
}

type CreatePlayListResponse struct {
	PlayListID int64 `json:"playlist_id"`
}

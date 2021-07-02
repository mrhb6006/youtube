package playListType

type CreatePlayListRequest struct {
	Name     string `json:"name" validate:"required"`
	IsPublic bool   `json:"is_public"`
}

type CreatePlayListResponse struct {
	PlayListID int64 `json:"playlist_id"`
}

type AddVideoToPlayListRequest struct {
	PlayListID int64 `json:"playlist_id" validate:"required"`
	VideoID    int64 `json:"video_id" validate:"required"`
}

type MakePublicRequest struct {
	PlayListID int64 `json:"playlist_id" validate:"required"`
}

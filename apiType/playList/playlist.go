package playList

type CreatePlayList struct {
	Name     string `json:"name"`
	IsPublic bool   `json:"is_public"`
}

type CreatePlayListResponse struct {
	PlayListID int64 `json:"playlist_id"`
}

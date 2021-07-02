package playlist

type PlayList struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	IsDefault bool   `json:"is_default"`
	IsPublic  bool   `json:"is_public"`
	CreatorID int64  `json:"creator_id"`
}

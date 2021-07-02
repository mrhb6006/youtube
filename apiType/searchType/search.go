package searchType

import "youtube/models/globalSearch"

type SearchRequest struct {
	Text string `json:"text"`
	Type string `json:"type" validate:"required,oneof=video playlist channel"`
}

type SearchResponse struct {
	Result []globalSearch.Search `json:"result"`
}

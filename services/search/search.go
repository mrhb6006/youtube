package search

import (
	"github.com/gofiber/fiber/v2"
	"strings"
	"youtube/apiType/commonType"
	"youtube/apiType/searchType"
	"youtube/models/globalSearch"
	"youtube/pkg/response"
)

func Search(ctx *fiber.Ctx) error {
	baseErrCode := "070"
	request := searchType.SearchRequest{}
	searchResponse := searchType.SearchResponse{}
	searchResponse.Result = []globalSearch.Search{}
	res := commonType.Response{}
	errStr, code, err := request.DecodeValidate(ctx)
	if err != nil {
		return response.ErrorResponse(ctx, res, baseErrCode, "01", errStr, code)
	}
	request.Text = strings.ToLower(strings.TrimSpace(request.Text))
	matched := make([]globalSearch.Search, 0)
	matched, errStr, err = globalSearch.Repo.Search(request.Text, request.Type)
	if err != nil {
		return response.ErrorResponse(ctx, res, baseErrCode, "02", errStr, 500)
	}
	searchResponse.Result = matched
	res.Res = searchResponse
	return response.SuccessResponse(ctx, res)
}

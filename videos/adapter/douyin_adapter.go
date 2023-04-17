package adapter

import (
	"github.com/assimon/svbot/internal/json"
	"github.com/assimon/svbot/internal/restyHttp"
)

type DouyinAdapter struct{}
type TenapiResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data struct {
	Title string `json:"title"`
	Author string `json:"author"`
	Avatar string `json:"avatar"`
	Like int `json:"like"`
	Time int64 `json:"time"`
	Cover string `json:"cover"`
	Url string `json:"url"`
	Music struct {
	Author string `json:"author"`
	Name string `json:"name"`
	} `json:"music"`
	} `json:"data"`
}

func (a DouyinAdapter) GetShortVideoInfo(url string) (*ShortVideoInfoResponse, error) {
	apiUrl := "https://tenapi.cn/v2/video?url=" + url
	resp, err := restyHttp.GetMobileHttpRequest().Get(apiUrl)
	if err != nil {
	return nil, err
	}
	var tenapiResponse TenapiResponse
	err = json.Cjson.Unmarshal(resp.Body(), &tenapiResponse)
	if err != nil {
	return nil, err
	}
	shortVideoInfo := &ShortVideoInfoResponse{
	AuthorName: tenapiResponse.Data.Author,
	AuthorAvatar: tenapiResponse.Data.Avatar,
	Title: tenapiResponse.Data.Title,
	Cover: tenapiResponse.Data.Cover,
	CreatedAt: "",
	NoWatermarkDownloadUrl: tenapiResponse.Data.Url,
	}
	return shortVideoInfo, nil
}

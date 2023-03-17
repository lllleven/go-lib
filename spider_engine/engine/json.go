package engine

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/tidwall/gjson"
)

func JsonHandler(ctx context.Context, s *Spider) ([]*SpiderData, error) {
	jsonResp, err := GetJsonResponse(ctx, s)
	if err != nil {
		return nil, err
	}
	datas := make([]*SpiderData, 0)
	if s.ListMatchRule != "" {
		listJson := gjson.Get(jsonResp, s.ListMatchRule)
		resultArray := listJson.Array()
		for _, result := range resultArray {
			data, _ := SetJsonDataMapping(ctx, s, result.String())
			datas = append(datas, data)
		}
	} else {
		data, _ := SetJsonDataMapping(ctx, s, jsonResp)
		datas = append(datas, data)
	}
	return datas, nil
}

func GetJsonResponse(ctx context.Context, s *Spider) (string, error) {
	client := http.Client{
		Timeout: time.Duration(s.TimeOut) * time.Second,
	}
	var reader io.Reader
	if s.Body != "" {
		reader = strings.NewReader(s.Body)
	}
	s.TargetUrl = SetHttpQuery(ctx, s.TargetUrl, s.Query)
	req, err := http.NewRequest(string(s.RequestMethod), s.TargetUrl, reader)
	if err != nil {
		return "", err
	}
	SetHttpHeader(ctx, req, s.Header)
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body), nil
}

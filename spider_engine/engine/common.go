package engine

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/simplejia/utils"
	"github.com/tidwall/gjson"
)

func SetHttpQuery(ctx context.Context, url string, queryMap map[string]string) string {
	if queryMap == nil {
		return url
	}
	// 如果strings.TrimRight() url最右边字符为/则去除
	url = strings.TrimRight(url, "/")
	// 如果url最右边字符不为?则添加?
	if !strings.HasSuffix(url, "?") {
		url += "?"
	}
	for k, v := range queryMap {
		url += k + "=" + v + "&"
	}
	return url[:len(url)-1]
}

func SetHttpHeader(ctx context.Context, req *http.Request, headerMap map[string]string) {
	for k, v := range headerMap {
		req.Header.Set(k, v)
	}
}

func SetJsonDataMapping(ctx context.Context, s *Spider, jsonString string) (*SpiderData, error) {
	data := new(SpiderData)
	for _, mapping := range s.FieldMappings {
		result := gjson.Get(jsonString, mapping.MatchRule)
		if result.String() != "" {
			SetSpiderData(ctx, data, mapping.Name, result.String())
		}
	}
	return data, nil
}

func SetSpiderData(ctx context.Context, data *SpiderData, key string, value string) error {
	if value == "" {
		return nil
	}
	switch key {
	case "id":
		data.Id = value
	case "title":
		data.Title = value
	case "content":
		data.Content = value
	case "desc":
		data.Desc = value
	case "url":
		data.Url = value
	case "cover":
		data.Cover = value
	default:
		if data.Extra != nil {
			data.Extra[key] = value
		} else {
			extra := map[string]string{}
			extra[key] = value
			data.Extra = extra
		}
	}
	return nil
}

func GetHtmlDocument(ctx context.Context, s *Spider) (*goquery.Document, error) {
	httpHeaderRet := make(http.Header)
	gpp := &utils.GPP{
		Uri:           s.TargetUrl,
		HttpHeaderRet: &httpHeaderRet,
		Timeout:       time.Duration(s.TimeOut) * time.Second,
		Headers:       s.Header,
	}
	body := make([]byte, 0)
	var err error
	switch s.RequestMethod {
	case RequestMethodGet:
		gpp.Params = s.Query
		body, err = utils.Get(gpp)
	case RequestMethodPost:
		gpp.Params = s.Body
		body, err = utils.Post(gpp)
	default:
		return nil, fmt.Errorf("request method %s is not support", s.RequestMethod)
	}
	if err != nil {
		return nil, err
	}
	reader := bytes.NewReader(body)
	return goquery.NewDocumentFromReader(reader)
}

func ParseHtmlMatchRule(ctx context.Context, matchRule string) (selector, method, attrName string, err error) {
	selectArray := strings.Split(matchRule, "=>")
	if len(selectArray) < 2 {
		return "", "", "", fmt.Errorf("match rule %s is not correct", matchRule)
	}
	selector = selectArray[0]
	method = selectArray[1]
	if method[:1] == ":" {
		attrName = method[1:]
		method = "attr"
	}
	return
}

func GetSelectDataByRule(ctx context.Context, method string, attrName string, selection *goquery.Selection) string {
	switch method {
	case "text":
		return selection.Text()
	case "html":
		html, _ := selection.Html()
		return html
	case "attr":
		attr, _ := selection.Attr(attrName)
		return attr
	}
	return ""
}

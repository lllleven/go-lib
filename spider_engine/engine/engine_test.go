package engine

import (
	"context"
	"fmt"
	"go-lib/util"
	"testing"
)

func TestJson(t *testing.T) {
	ctx := context.Background()
	s := Spider{
		Name: "test",
		Origin: map[string]string{
			"origin_test": "origin_test",
		},
		TargetUrl:    "https://api.bitea.one/post/list_att?sign=18d7c76999e9b8b69c18b77a4958da24",
		ResponseType: ResponseTypeJson,
		Header: map[string]string{
			"debug": "1",
		},
		Body:          `{"h_av":"0.2.2-nightly","h_dt":0,"h_os":"30","h_app":"bitea","h_model":"M2012K11AC","h_ch":"others","h_did":"5823886f-3f33-4afd-9a4d-be6a8057f591","h_m":11983580,"token":"TcKbN2AH_JcPkM-igo3EJjSAbIlTkew7VlsJ4XTuChnNwrNXpjRIcT70eJdfmvBcFO9sa20ydtIaDwoP6SDsTQH0W64v6tCRyKdOHVAvQQZWi848=","h_nt":1,"h_ts":1659080094467}`,
		RequestMethod: RequestMethodPost,
		ListMatchRule: "data.list",
		FieldMappings: []*FieldMapping{
			{
				Name:      "content",
				MatchRule: "content",
			},
			{
				Name:      "avatar",
				MatchRule: "member.avatar",
			},
		},
	}
	dataList, err := s.Exec(ctx)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(util.JsonPretty(dataList))
}

func TestHtml(t *testing.T) {
	ctx := context.Background()
	s := Spider{
		Name: "test",
		Origin: map[string]string{
			"spider": "36kr",
		},
		TargetUrl:    "https://36kr.com/information/web_news/",
		ResponseType: ResponseTypeHtml,
		Header: map[string]string{
			"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:103.0) Gecko/20100101 Firefox/103.0",
		},
		RequestMethod: RequestMethodGet,
		ListMatchRule: ".information-flow-list .information-flow-item",
		FieldMappings: []*FieldMapping{
			{
				Name:      "title",
				MatchRule: ".kr-shadow-content a[class='article-item-title weight-bold']=>text",
			},
			{
				Name:      "url",
				MatchRule: "a[class='article-item-description ellipsis-2']=>:href",
			},
		},
	}
	dataList, err := s.Exec(ctx)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(util.JsonPretty(dataList))
}

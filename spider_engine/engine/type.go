package engine

type Spider struct {
	Id            int64             `json:"id" bson:"_id"`
	Name          string            `json:"name" bson:"name"`
	TargetUrl     string            `json:"target_url" bson:"target_url"`
	RequestMethod RequestMethod     `json:"request_method" bson:"request_method"`
	Header        map[string]string `json:"header" bson:"header"`
	Body          string            `json:"body" bson:"body"`
	Query         map[string]string `json:"query" bson:"query"`
	TimeOut       int64             `json:"timeout" bson:"timeout"`
	ResponseType  ResponseType      `json:"response_type" bson:"response_type"`
	ListMatchRule string            `json:"list_match_rule" bson:"list_match_rule"`
	// 当 ResponseType == ResponseTypeHtmlToJson 有效 Html中获取json结构
	JsonMatchRule string          `json:"json_match_rule" bson:"json_match_rule"`
	FieldMappings []*FieldMapping `json:"field_mappings" bson:"field_mappings"`
	// 透传
	Origin map[string]string `json:"origin" bson:"origin"`
}

type FieldMapping struct {
	Name string `json:"name" bson:"name"`
	// 匹配规则
	MatchRule string `json:"match_rule" bson:"match_rule"`
}

type SpiderData struct {
	Id      string            `json:"id"`
	Url     string            `json:"url"`
	Title   string            `json:"title"`
	Content string            `json:"content"`
	Cover   string            `json:"cover"`
	Desc    string            `json:"desc"`
	Extra   map[string]string `json:"extra"`
	Origin  map[string]string `json:"origin"`
}

type RequestMethod string

const (
	RequestMethodGet  RequestMethod = "GET"
	RequestMethodPost RequestMethod = "POST"
)

type ResponseType string

const (
	ResponseTypeHtml       ResponseType = "HTML"
	ResponseTypeJson       ResponseType = "JSON"
	ResponseTypeHtmlToJson ResponseType = "HTMLTOJSON"
)

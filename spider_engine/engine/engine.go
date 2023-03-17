package engine

import (
	"context"
	"fmt"
)

func (s *Spider) Exec(ctx context.Context) ([]*SpiderData, error) {
	if err := s.SelfCheck(); err != nil {
		return nil, err
	}
	dataList := make([]*SpiderData, 0)
	var err error
	switch s.ResponseType {
	case ResponseTypeJson:
		dataList, err = JsonHandler(ctx, s)
		if err != nil {
			return nil, err
		}
		for _, data := range dataList {
			data.Origin = s.Origin
		}
	case ResponseTypeHtml:
		dataList, err = HtmlHandler(ctx, s)
		if err != nil {
			return nil, err
		}
		for _, data := range dataList {
			data.Origin = s.Origin
		}
	case ResponseTypeHtmlToJson:

	default:
		return nil, fmt.Errorf("ResponseType: %v is not support", s.ResponseType)
	}
	return dataList, nil
}

func (s *Spider) SelfCheck() error {
	if s.ResponseType == "" {
		return fmt.Errorf("ResponseType is empty")
	}
	if s.TargetUrl == "" {
		return fmt.Errorf("TargetUrl is empty")
	}
	if s.TimeOut == 0 {
		s.TimeOut = 30
	}
	return nil
}

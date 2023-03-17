package engine

import (
	"context"

	"github.com/PuerkitoBio/goquery"
)

func HtmlHandler(ctx context.Context, s *Spider) ([]*SpiderData, error) {
	document, err := GetHtmlDocument(ctx, s)
	if err != nil {
		return nil, err
	}
	selection := document.Selection
	datas := make([]*SpiderData, 0)
	if s.ListMatchRule != "" {
		selection.Find(s.ListMatchRule).Each(func(i int, items *goquery.Selection) {
			data, _ := SetHtmlDataMapping(ctx, s, items)
			datas = append(datas, data)
		})
	} else {
		data, _ := SetHtmlDataMapping(ctx, s, selection)
		datas = append(datas, data)
	}
	return datas, nil
}

func SetHtmlDataMapping(ctx context.Context, stage *Spider, selection *goquery.Selection) (*SpiderData, error) {
	data := new(SpiderData)
	for _, mapping := range stage.FieldMappings {
		selector, method, attrName, _ := ParseHtmlMatchRule(ctx, mapping.MatchRule)
		if selector == "" {
			value := GetSelectDataByRule(ctx, method, attrName, selection)
			if value != "" {
				SetSpiderData(ctx, data, mapping.Name, value)
			}
		} else {
			selection.Find(selector).Each(func(i int, selection *goquery.Selection) {
				value := GetSelectDataByRule(ctx, method, attrName, selection)
				if value != "" {
					SetSpiderData(ctx, data, mapping.Name, value)
				}
			})
		}
	}
	return data, nil
}

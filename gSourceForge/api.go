package gSourceForge

import (
	"fmt"
	"github.com/gek64/gek/gXml"
	"strings"
)

type API struct {
	Channel Channel `xml:"channel"`
}

type Channel struct {
	Title string `xml:"title"`
	Item  []Item `xml:"item"`
}

type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Guid        string `xml:"guid"`
	PubDate     string `xml:"pubDate"`
	Description string `xml:"description"`
}

// NewAPI 新建 API
func NewAPI(rssUrl string) (api *API, err error) {
	// 新建xml处理体
	xmlOperator, err := gXml.NewXmlOperator(api)
	if err != nil {
		return nil, err
	}

	// xml处理体从URL中读取xml数据,数据存储到api中
	err = xmlOperator.ReadFromURL(rssUrl)
	if err != nil {
		return nil, err
	}

	return api, nil
}

// SearchPartsInRelease 搜索 API 中 item/title 中的多个名称,返回第一个全匹配的下载链接
func (api API) SearchPartsInRelease(parts []string) (downloadUrl string, err error) {

	for _, i := range api.Channel.Item {
		matched := true

		for _, part := range parts {
			if !strings.Contains(i.Title, part) {
				matched = false
				break
			}
		}

		if matched {
			return i.Link, nil
		}
	}

	return "", fmt.Errorf("can not find release with parts %s", parts)
}

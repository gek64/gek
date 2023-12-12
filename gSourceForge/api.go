package gSourceForge

import (
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
	xmlOperator, err := gXml.NewXmlOperator(&api)
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

// SearchRelease 搜索
func (api API) SearchRelease(includes []string, excludes []string) (items []Item) {
	var list []Item

	// 排除不包含
	for _, exclude := range excludes {
		for _, item := range api.Channel.Item {
			if !strings.Contains(item.Title, exclude) {
				list = append(list, item)
			}
		}
	}
	// 寻找所有全包含项目
	for i, item := range list {
		var matched = true
		for _, include := range includes {
			if !strings.Contains(item.Title, include) {
				matched = false
				break
			}
		}
		if matched {
			items = append(items, list[i])
		}
	}
	return items
}

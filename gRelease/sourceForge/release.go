package sourceForge

import (
	"github.com/gek64/gek/gXml"
	"slices"
	"strings"
)

type Release struct {
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

func NewRelease(releaseApiUrl string) (r *Release, err error) {
	// 新建xml处理体
	xmlOperator, err := gXml.NewXmlOperator(&r)
	if err != nil {
		return nil, err
	}
	// xml处理体从URL中读取xml数据,数据存储到api中
	err = xmlOperator.ReadFromURL(releaseApiUrl)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// SearchRelease 搜索
func (r *Release) SearchRelease(includes []string, excludes []string) (assets []Item) {
	// 排除不包含
	for _, exclude := range excludes {
		r.Channel.Item = slices.DeleteFunc(r.Channel.Item, func(assets Item) bool {
			return strings.Contains(assets.Title, exclude)
		})
	}

	// 寻找所有全包含项目
	for _, include := range includes {
		r.Channel.Item = slices.DeleteFunc(r.Channel.Item, func(assets Item) bool {
			return !strings.Contains(assets.Title, include)
		})
	}

	return r.Channel.Item
}

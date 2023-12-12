package gGithub

import (
	"github.com/gek64/gek/gJson"
	"strings"
)

// API Github 下载 API
type API struct {
	TagName    string   `json:"tag_name"`
	Assets     []Assets `json:"assets"`
	TarballURL string   `json:"tarball_url"`
	ZipballURL string   `json:"zipball_url"`
	Body       string   `json:"body"`
}
type Assets struct {
	Name               string `json:"name"`
	ContentType        string `json:"content_type"`
	BrowserDownloadURL string `json:"browser_download_url"`
}

// NewAPI 新建 API
func NewAPI(repo string) (api *API, err error) {
	// 新建json处理体
	jsonOperator, err := gJson.NewJsonOperator(&api)
	if err != nil {
		return nil, err
	}

	// json处理体从URL中读取json数据,数据存储到githubAPI中
	err = jsonOperator.ReadFromURL("https://api.github.com/repos/" + repo + "/releases/latest")
	if err != nil {
		return nil, err
	}

	return api, nil
}

// SearchRelease 搜索
func (api API) SearchRelease(includes []string, excludes []string) (assets []Assets) {
	var list []Assets

	// 排除不包含
	for _, exclude := range excludes {
		for _, asset := range api.Assets {
			if !strings.Contains(asset.Name, exclude) {
				list = append(list, asset)
			}
		}
	}
	// 寻找所有全包含项目
	for i, asset := range list {
		var matched = true
		for _, include := range includes {
			if !strings.Contains(asset.Name, include) {
				matched = false
				break
			}
		}
		if matched {
			assets = append(assets, list[i])
		}
	}
	return assets
}

package github

import (
	"github.com/gek64/gek/gJson"
	"slices"
	"strings"
)

type Release struct {
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

// NewRelease 新建 Release
// https://docs.github.com/en/rest/releases
func NewRelease(releaseApiUrl string) (r *Release, err error) {
	// 新建json处理体
	jsonOperator, err := gJson.NewJsonOperator(&r)
	if err != nil {
		return nil, err
	}

	// json处理体从URL中读取json数据,数据存储到githubAPI中
	err = jsonOperator.ReadFromURL(releaseApiUrl)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func NewReleaseLatest(repo string) (r *Release, err error) {
	return NewRelease("https://api.github.com/repos/" + repo + "/releases/latest")
}

func NewReleaseByTag(repo string, tag string) (r *Release, err error) {
	return NewRelease("https://api.github.com/repos/" + repo + "/releases/tags/" + tag)
}

func NewReleaseById(repo string, id string) (r *Release, err error) {
	return NewRelease("https://api.github.com/repos/" + repo + "/releases/" + id)
}

// SearchRelease 搜索
func (r *Release) SearchRelease(includes []string, excludes []string) (assets []Assets) {
	// 排除不包含
	for _, exclude := range excludes {
		r.Assets = slices.DeleteFunc(r.Assets, func(assets Assets) bool {
			return strings.Contains(assets.Name, exclude)
		})
	}
	// 寻找所有全包含项目
	for _, include := range includes {
		r.Assets = slices.DeleteFunc(r.Assets, func(assets Assets) bool {
			return !strings.Contains(assets.Name, include)
		})
	}
	return r.Assets
}

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

// https://docs.github.com/en/rest/releases
// https://api.github.com/repos/microsoft/terminal/releases
func newRelease[T *[]Release | *Release](releaseApiUrl string) (r T, err error) {
	// 新建 json 处理器
	jsonOperator, err := gJson.NewJsonOperator(&r)
	if err != nil {
		return nil, err
	}
	// 从 releaseApiUrl 中读取数据存储到结构体中
	return r, jsonOperator.ReadFromURL(releaseApiUrl)
}

func GetReleases(repo string) (rs *[]Release, err error) {
	return newRelease[*[]Release]("https://api.github.com/repos/" + repo + "/releases")
}

func GetReleaseLatest(repo string) (r *Release, err error) {
	return newRelease[*Release]("https://api.github.com/repos/" + repo + "/releases/latest")
}

func GetReleaseByTagName(repo string, tagName string) (r *Release, err error) {
	return newRelease[*Release]("https://api.github.com/repos/" + repo + "/releases/tags/" + tagName)
}

func (r *Release) GetAssets(includes []string, excludes []string) (assets []Assets) {
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

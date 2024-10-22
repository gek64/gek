package gitlab

import (
	"github.com/gek64/gek/gJson"
	"slices"
	"strings"
)

type Release struct {
	Name            string `json:"name"`
	TagName         string `json:"tag_name"`
	UpcomingRelease bool   `json:"upcoming_release"`
	Assets          Assets `json:"assets"`
}
type Assets struct {
	Count   int       `json:"count"`
	Sources []Sources `json:"sources"`
	Links   []Links   `json:"links"`
}
type Sources struct {
	Format string `json:"format"`
	URL    string `json:"url"`
}
type Links struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	URL            string `json:"url"`
	DirectAssetURL string `json:"direct_asset_url"`
	LinkType       string `json:"link_type"`
}

// https://docs.gitlab.com/ee/api/releases
// https://gitlab.com/api/v4/projects/36189/releases
// https://gitlab.com/api/v4/projects/fdroid%2ffdroidclient/releases
func newRelease[T *[]Release | *Release](releaseApiUrl string) (r T, err error) {
	// 新建 json 处理器
	jsonOperator, err := gJson.NewJsonOperator(&r)
	if err != nil {
		return nil, err
	}
	// 从 releaseApiUrl 中读取数据存储到结构体中
	return r, jsonOperator.ReadFromURL(releaseApiUrl)
}

func GetReleases(projectId string) (rs *[]Release, err error) {
	return newRelease[*[]Release]("https://gitlab.com/api/v4/projects/" + projectId + "/releases")
}

func GetReleaseLatest(projectId string) (r *Release, err error) {
	return newRelease[*Release]("https://gitlab.com/api/v4/projects/" + projectId + "/releases/permalink/latest")
}

func GetReleaseByTagName(projectId string, tagName string) (r *Release, err error) {
	return newRelease[*Release]("https://gitlab.com/api/v4/projects/" + projectId + "/releases/" + tagName)
}

func (r *Release) GetAssets(includes []string, excludes []string) (l []Links) {
	// 排除不包含
	for _, exclude := range excludes {
		r.Assets.Links = slices.DeleteFunc(r.Assets.Links, func(l Links) bool {
			return strings.Contains(l.Name, exclude)
		})
	}
	// 寻找所有全包含项目
	for _, include := range includes {
		r.Assets.Links = slices.DeleteFunc(r.Assets.Links, func(l Links) bool {
			return !strings.Contains(l.Name, include)
		})
	}
	return r.Assets.Links
}

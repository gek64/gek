package gGithub

import (
    "fmt"
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

// SearchRelease 搜索 API 中 Assets 中的名称,返回第一个匹配的下载链接
func (api API) SearchRelease(part string) (downloadUrl string, err error) {
    for _, asset := range api.Assets {
        if strings.Contains(asset.Name, part) {
            return asset.BrowserDownloadURL, nil
        }
    }
    return "", fmt.Errorf("can not find release with part %s", part)
}

// SearchPartsInRelease 搜索 API 中 Assets 中的多个名称,返回第一个匹配的下载链接
func (api API) SearchPartsInRelease(parts []string) (downloadUrl string, err error) {

    for _, asset := range api.Assets {
        var matched bool = true

        for _, part := range parts {
            if !strings.Contains(asset.Name, part) {
                matched = false
                break
            }
        }

        if matched {
            return asset.BrowserDownloadURL, nil
        }
    }

    return "", fmt.Errorf("can not find release with parts %s", parts)
}

package deprecated

import (
	"container/list"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// WildcardToRegex 通配符 * ? 转换为正则表达式
func WildcardToRegex(wildcard string) string {
	if strings.Contains(wildcard, "*") || strings.Contains(wildcard, "?") {
		regex := strings.ReplaceAll(wildcard, "*", ".*")
		regex = strings.ReplaceAll(regex, "?", ".?")
		return regex
	}
	return wildcard
}

// MatchString 字符串匹配字符串
func MatchString(pattern string, s string) bool {
	matched := false

	// 如果pattern是正则表达式的话,使用正则表达式匹配
	_, err := regexp.Compile(pattern)
	if err == nil {
		matched, err = regexp.MatchString(pattern, s)
	}

	// 如果不是正则表达式,就进行字符串匹配
	if err != nil {
		matched = strings.Contains(s, pattern)
	}

	return matched
}

// MatchFilter 过滤表匹配字符串
func MatchFilter(filter []string, s string) bool {
	matched := false

	// 循环匹配过滤表内的pattern
	for _, pattern := range filter {
		matched = MatchString(pattern, s)
		if matched {
			break
		}
	}
	return matched
}

// WalkAll 递归将指定路径下所有文件加入列表(链表或者切片)
// path为指定的文件路径字符串,listPointer为加入列表的指针,subFolder为是否要递归下级目录的布尔值
func WalkAll(path string, listPointer any, subFolder bool) error {
	// 读取文件下所有文件以及文件夹
	fileInfo, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for _, file := range fileInfo {
		// 获取文件的绝对路径
		absPath, err := filepath.Abs(path)
		if err != nil {
			return err
		}
		itemPath := filepath.Join(absPath, file.Name())

		// 如果file是文件,加入到列表当中
		if !file.IsDir() {
			switch listPointer.(type) {
			case *[]string:
				l := listPointer.(*[]string)
				*l = append(*l, itemPath)
			case *list.List:
				l := listPointer.(*list.List)
				l.PushBack(itemPath)
			default:
				return fmt.Errorf("wrong type of input list")
			}
		}

		// 如果file是目录并且指定递归处理,递归处理此目录下文件
		if file.IsDir() && subFolder {
			err := WalkAll(itemPath, listPointer, true)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

package gek_file

import (
	"container/list"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var OriginalWorkingPath, _ = os.Getwd()

// PathCreate Create path if the path not exists
func PathCreate(path string) bool {
	if !PathExist(path) {
		err := os.MkdirAll(path, 700)
		log.Println(err)
		return true
	}
	return false
}

// PathWalk 递归将指定路径下文件加入列表(链表或者切片),带过滤器
func PathWalk(path string, listPointer interface{}, filter []string, subFolder bool) error {
	// 读取文件下所有文件以及文件夹
	fileInfo, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	for _, file := range fileInfo {
		// 获取文件的绝对路径,处理"\"
		itemPath := strings.Replace(filepath.Join(path, file.Name()), "\\", "/", -1)

		// 当过滤表不为空时,进行过滤操作
		if len(filter) != 0 {
			matched := false
			matched = MatchFilter(filter, itemPath)
			// 如果符合过滤列表则跳过当前文件,并继续处理下一个文件
			if matched {
				fmt.Printf("skip:%s\n", itemPath)
				continue
			}
		}

		// 如果当前工作路径不是函数输入路径,切换到函数输入路径
		if gwd, _ := os.Getwd(); gwd != path {
			err := os.Chdir(path)
			if err != nil {
				return err
			}
		}

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
				return fmt.Errorf("wrong type of input l")
			}
		}

		// 如果file是目录并且指定递归处理,递归处理此目录下文件
		if file.IsDir() && subFolder {
			err := PathWalk(itemPath, listPointer, filter, true)
			if err != nil {
				return err
			}
		}
	}

	// 切换回到原有工作路径
	err = os.Chdir(OriginalWorkingPath)
	if err != nil {
		return err
	}
	return nil
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

// PathExist Check if path string exists
func PathExist(path string) bool {
	var _, err = os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

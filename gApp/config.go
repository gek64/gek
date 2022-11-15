package gApp

import (
	"fmt"
	"github.com/gek64/gek/gFile"
	"os"
	"path/filepath"
)

type Config struct {
	// 配置文件名称
	Name string
	// 配置文件内容
	Content string
	// 配置文件安装文件夹
	Location string
	// 删除时是否删除安装路径
	UninstallDeleteLocationFolder bool
}

// NewConfig 新建配置文件
func NewConfig(name string, content string, location string, uninstallDeleteLocationFolder bool) (c Config) {
	return Config{Name: name, Content: content, Location: location, UninstallDeleteLocationFolder: uninstallDeleteLocationFolder}
}

// NewConfigFromFile 新建配置文件(从文件)
func NewConfigFromFile(name string, file string, uninstallDeleteLocationFolder bool, location string) (c Config, err error) {
	bytes, err := os.ReadFile(file)
	if err != nil {
		return Config{}, err
	}
	return NewConfig(name, string(bytes), location, uninstallDeleteLocationFolder), nil
}

// Install 安装配置文件
func (c Config) Install() (err error) {
	// 检测配置安装路径是否存在
	// 不存在则创建
	_, err = os.Stat(c.Location)
	if os.IsNotExist(err) {
		err = os.MkdirAll(c.Location, 00755)
		if err != nil {
			return err
		}
	}

	_, err = gFile.CreateFile(filepath.Join(c.Location, c.Name), c.Content)
	if err != nil {
		return err
	}

	// 赋权0644
	err = os.Chmod(filepath.Join(c.Location, c.Name), 0644)

	return err
}

// Uninstall 卸载配置文件,并删除配置文件安装文件夹
func (c Config) Uninstall() (err error) {
	// 检测配置安装路径是否存在
	_, err = os.Stat(c.Location)
	if os.IsNotExist(err) {
		return fmt.Errorf("can't find config location %s", c.Location)
	}

	if c.UninstallDeleteLocationFolder {
		err = os.RemoveAll(c.Location)
	} else {
		err = os.RemoveAll(filepath.Join(c.Location, c.Name))
	}
	return err
}

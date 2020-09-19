package gopkg

import (
	"encoding/json"
	"io/ioutil"
)

// ReadJSON 输入结构体指针与json文件路径,将json内部数据存储到结构体中
func ReadJSON(stPointer interface{}, jsonFile string) error {
	//打开json文件
	fileData, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		return err
	}
	//json数据写入结构体
	err = json.Unmarshal(fileData, stPointer)
	if err != nil {
		return err
	}
	return nil
}

// WriteJSON 写入json文件
func WriteJSON(st interface{}, fileName string) error {
	//从结构体生成json的byte数据
	jsonData, err := json.MarshalIndent(st, "", "    ")
	if err != nil {
		return err
	}
	//json的byte数据写入json文件
	err = ioutil.WriteFile(fileName, jsonData, 0644)
	if err != nil {
		return err
	}
	return nil
}

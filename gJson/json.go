package gJson

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"reflect"
)

// JsonOperator json处理体
type JsonOperator struct {
	jsonStructPointer any
}

// NewJsonOperator 输入结构体指针传入json处理体,jsonStructPointer是需要存储到的结构体实例的指针
func NewJsonOperator(jsonStructPointer any) (*JsonOperator, error) {
	// 判断传入v是否是指针类型
	if reflect.ValueOf(jsonStructPointer).Kind() != reflect.Ptr {
		return nil, fmt.Errorf("%v is not of pointer type", jsonStructPointer)
	}
	// 填充json处理体
	j := JsonOperator{
		jsonStructPointer: jsonStructPointer,
	}
	return &j, nil
}

// ReadFromFile 输入结构体指针与json文件路径,将json文件数据存储到json处理体中
func (j *JsonOperator) ReadFromFile(filename string) error {
	// 打开json文件
	fileData, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	// json数据写入结构体
	err = json.Unmarshal(fileData, j.jsonStructPointer)
	if err != nil {
		return err
	}
	return nil
}

// ReadFromURL 输入json文件URL,将url中包含的json数据存储到json处理体中
func (j *JsonOperator) ReadFromURL(url string) error {
	// 打开链接
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	// 将response中的json数据解析，然后写入处理体
	err = json.NewDecoder(response.Body).Decode(j.jsonStructPointer)
	if err != nil {
		return err
	}
	return nil
}

// WriteToFile 写入json文件，filename是要写入的文件的名称
func (j *JsonOperator) WriteToFile(filename string) error {
	// 从结构体生成json的byte数据
	jsonData, err := json.Marshal(j.jsonStructPointer)
	if err != nil {
		return err
	}
	// json的byte数据写入json文件
	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return err
	}
	return nil
}

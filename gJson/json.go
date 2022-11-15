package gJson

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"reflect"
)

// Jsoner json处理体
type Jsoner struct {
	Jst interface{}
}

// NewJsoner 输入结构体指针传入处理体jsoner,v是需要存储到的结构体实例的指针
func NewJsoner(v interface{}) (*Jsoner, error) {
	// 判断传入v是否是指针类型
	if reflect.ValueOf(v).Kind() != reflect.Ptr {
		return nil, fmt.Errorf("%v is not of pointer type", v)
	}
	// 填充json处理体
	j := Jsoner{
		Jst: v,
	}
	return &j, nil
}

// ReadFromFile 输入结构体指针与json文件路径,将json文件数据存储到json处理体中
func (j *Jsoner) ReadFromFile(filename string) error {
	// 打开json文件
	fileData, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	// json数据写入结构体
	err = json.Unmarshal(fileData, j.Jst)
	if err != nil {
		return err
	}
	return nil
}

// ReadFromURL 输入json文件URL,将url中包含的json数据存储到json处理体中
func (j *Jsoner) ReadFromURL(url string) error {
	// 打开链接
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	// 将response中的json数据解析，然后写入处理体
	err = json.NewDecoder(response.Body).Decode(&j.Jst)
	if err != nil {
		return err
	}
	return nil
}

// WriteToFile 写入json文件，filename是要写入的文件的名称
func (j *Jsoner) WriteToFile(filename string) error {
	// 从结构体生成json的byte数据
	jsonData, err := json.MarshalIndent(j.Jst, "", "    ")
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

//// ReadFromFile 输入结构体指针与json文件路径,将json内部数据存储到结构体中,v是需要存储到的结构体实例的指针
//func ReadFromFile(v interface{}, filename string) error {
//	//打开json文件
//	fileData, err := ioutil.ReadFile(filename)
//	if err != nil {
//		return err
//	}
//	//json数据写入结构体
//	err = json.Unmarshal(fileData, v)
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//// WriteToFile 写入json文件,v是需要存储的结构体实例的指针
//func WriteToFile(v interface{}, filename string) error {
//	//从结构体生成json的byte数据
//	jsonData, err := json.MarshalIndent(v, "", "    ")
//	if err != nil {
//		return err
//	}
//	//json的byte数据写入json文件
//	err = ioutil.WriteFile(filename, jsonData, 0644)
//	if err != nil {
//		return err
//	}
//	return nil
//}

package gek_json

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
)

// Jsoner json处理体
type Jsoner struct {
	Jst interface{}
}

// NewJsoner 输入结构体指针传入处理体jsoner,v是需要存储到的结构体实例的指针
func NewJsoner(v interface{}) (*Jsoner, error) {
	if reflect.ValueOf(v).Kind() != reflect.Ptr {
		return nil, fmt.Errorf("%v is not of pointer type", v)
	}
	j := Jsoner{
		Jst: v,
	}
	return &j, nil
}

// ReadFromFile 输入结构体指针与json文件路径,将json内部数据存储到结构体中,v是需要存储到的结构体实例的指针
func (j *Jsoner) ReadFromFile(filename string) error {
	//打开json文件
	fileData, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	//json数据写入结构体
	err = json.Unmarshal(fileData, j.Jst)
	if err != nil {
		return err
	}
	return nil
}

// WriteToFile 写入json文件,v是需要存储的结构体实例的指针
func (j *Jsoner) WriteToFile(filename string) error {
	//从结构体生成json的byte数据
	jsonData, err := json.MarshalIndent(j.Jst, "", "    ")
	if err != nil {
		return err
	}
	//json的byte数据写入json文件
	err = ioutil.WriteFile(filename, jsonData, 0644)
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

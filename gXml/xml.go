package gXml

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"os"
	"reflect"
)

// XmlOperator xml处理体
type XmlOperator struct {
	xmlStructPointer any
}

// NewXmlOperator 新建xml处理体, xmlPointer是需要存储到的结构体实例的指针
func NewXmlOperator(xmlStructPointer any) (*XmlOperator, error) {
	// 判断传入v是否是指针类型
	if reflect.ValueOf(xmlStructPointer).Kind() != reflect.Ptr {
		return nil, fmt.Errorf("%v is not of pointer type", xmlStructPointer)
	}
	// 填充xml处理体
	return &XmlOperator{xmlStructPointer: xmlStructPointer}, nil
}

// ReadFromFile 输入结构体指针与xml文件路径,将xml文件数据存储到xml处理体中
func (x *XmlOperator) ReadFromFile(filename string) error {
	// 打开文件
	file, err := os.Open(filename)
	// xml数据解析后,写入结构体
	err = xml.NewDecoder(file).Decode(x.xmlStructPointer)
	if err != nil {
		return err
	}
	return nil
}

// ReadFromURL 输入xml文件URL,将url中包含的xml数据解析后,存储到xml处理体中
func (x *XmlOperator) ReadFromURL(url string) error {
	// 打开链接
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	// 将response中的xml数据解析，然后写入处理体
	err = xml.NewDecoder(response.Body).Decode(x.xmlStructPointer)
	if err != nil {
		return err
	}
	return nil
}

// WriteToFile 写入xml文件，filename是要写入的文件的名称
func (x *XmlOperator) WriteToFile(filename string) error {
	// 从结构体生成xml的byte数据
	xmlData, err := xml.Marshal(x.xmlStructPointer)
	if err != nil {
		return err
	}
	// xml的byte数据写入文件
	err = os.WriteFile(filename, xmlData, 0644)
	if err != nil {
		return err
	}
	return nil
}

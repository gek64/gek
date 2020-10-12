package gopkg

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type testJson struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func ExampleJsoner_ReadFromFile() {
	// 创建临时json文件
	f, err := CreateTempFile("./", "test*.json", "{\"name\":\"bob\",\"age\":24}")
	if err != nil {
		log.Fatal(err)
	}
	// 删除临时json文件
	defer os.Remove(f.Name())

	// 新建结构体实例
	tj := new(testJson)

	// 使用json处理体读取json文件到结构体实例中
	jsoner := new(Jsoner)
	err = jsoner.New(&tj)
	if err != nil {
		log.Fatal(err)
	}
	err = jsoner.ReadFromFile(f.Name())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(*tj)

	// Output:
	// {bob 24}
}

func ExampleJsoner_WriteToFile() {
	// 创建结构体实例
	var tj = testJson{
		Name: "bob",
		Age:  24,
	}

	jsoner := new(Jsoner)
	err := jsoner.New(&tj)
	if err != nil {
		log.Fatal(err)
	}
	err = jsoner.WriteToFile("test.json")

	defer os.Remove("test.json")

	text, err := ioutil.ReadFile("test.json")

	fmt.Print(string(text), err)

	// Output:
	// {
	//     "name": "bob",
	//     "age": 24
	// }<nil>

}

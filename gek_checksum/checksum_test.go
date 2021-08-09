package gek_checksum

import (
	"gek_file"
	"os"
	"testing"
)

func TestChecksum(t *testing.T) {

	// 创建临时文件
	tmpFile, err := gek_file.CreateTempFile("./", "testFile*.txt", "This is a temporary file!")
	if err != nil {
		t.Fatal(err)
	}
	// 销毁时删除临时文件
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			t.Fatal(err)
		}
	}(tmpFile.Name())

	// Table-Driven 测试用例
	var checksumTests = []struct {
		mode     string
		fileURL  string
		expected string
	}{
		{"crc32", tmpFile.Name(), "d63e2c3f"},
		{"md5", tmpFile.Name(), "48851fd844cbf431b4393158b984524f"},
		{"sha1", tmpFile.Name(), "b73697f59d2e5cd702a36ad89c3f910d404757b6"},
		{"sha256", tmpFile.Name(), "ecb81149c0f578af6115e847c7b6f20b86aabea9ad4ba00b3f795d9bf4455ab2"},
	}

	// 测试核心
	for _, tt := range checksumTests {
		actual := Checksum(tt.mode, tt.fileURL)
		if actual != tt.expected {
			t.Errorf("Checksum(%s, %s) = %s; expected %s", tt.mode, tt.fileURL, actual, tt.expected)
		}
	}

}

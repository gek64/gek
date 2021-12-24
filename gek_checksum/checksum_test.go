package gek_checksum

import (
	"os"
	"testing"
)

func TestChecksum(t *testing.T) {
	// 创建临时文件
	tmpFile, err := os.Create("tmpFile.txt")
	if err != nil {
		t.Fatal(err)
	}
	// 临时文件写入信息
	_, err = tmpFile.Write([]byte("This is a temporary file!"))
	if err != nil {
		t.Fatal(err)
	}
	// 销毁时关闭文件,删除临时文件
	defer func(tmpFile *os.File) {
		err := tmpFile.Close()
		if err != nil {
			t.Fatal(err)
		}
		err = os.Remove(tmpFile.Name())
		if err != nil {
			t.Fatal(err)
		}
	}(tmpFile)

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
		actual, err := Checksum(tt.mode, tt.fileURL)
		if err != nil {
			t.Errorf("%v", err)
		}
		if actual != tt.expected {
			t.Errorf("Checksum(%s, %s) = %s; expected %s", tt.mode, tt.fileURL, actual, tt.expected)
		}
	}

}

// 基准测试核心
func benchmarkChecksum(mode string, b *testing.B) {
	// 创建临时文件
	tmpFile, err := os.Create("tmpFile.txt")
	if err != nil {
		b.Fatal(err)
	}
	// 创建临时文件大小10MB
	err = tmpFile.Truncate(1e7)
	if err != nil {
		b.Fatal(err)
	}
	// 销毁时关闭文件,删除临时文件
	defer func(tmpFile *os.File) {
		err := tmpFile.Close()
		if err != nil {
			b.Fatal(err)
		}
		err = os.Remove(tmpFile.Name())
		if err != nil {
			b.Fatal(err)
		}
	}(tmpFile)

	// 重置基准测试计时器,来忽略初始化生成测试文件的时间
	b.ResetTimer()

	// 基准测试运行主函数
	for i := 0; i < b.N; i++ {
		_, err := Checksum(mode, tmpFile.Name())
		if err != nil {
			b.Fatal(err)
		}
	}
}

// 基准测试CRC32
func BenchmarkChecksumCrc32(b *testing.B) {
	benchmarkChecksum("crc32", b)
}

// 基准测试SHA1
func BenchmarkChecksumSha1(b *testing.B) {
	benchmarkChecksum("sha1", b)
}

// 基准测试SHA256
func BenchmarkChecksumSha256(b *testing.B) {
	benchmarkChecksum("sha256", b)
}

// 基准测试MD5
func BenchmarkChecksumMd5(b *testing.B) {
	benchmarkChecksum("md5", b)
}

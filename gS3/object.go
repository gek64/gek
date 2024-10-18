package gS3

import (
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"os"
)

// 用法 https://docs.aws.amazon.com/zh_cn/code-library/latest/ug/go_2_s3_code_examples.html

func (c *Client) UploadObject(bucket string, filename string, data []byte) (*manager.UploadOutput, error) {
	return manager.NewUploader(c).Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
		Body:   bytes.NewReader(data)},
	)
}

func (c *Client) DownloadObject(bucket string, filename string, downloadFilename string) (int64, error) {
	f, err := os.Create(downloadFilename)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	return manager.NewDownloader(c).Download(context.TODO(), f, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
	})
}

func (c *Client) DeleteObject(bucket string, filename string) (*s3.DeleteObjectOutput, error) {
	return c.Client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{Bucket: aws.String(bucket), Key: aws.String(filename)})
}

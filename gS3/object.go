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
	return manager.NewUploader(c.S3Client).Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
		Body:   bytes.NewReader(data)},
	)
}

func (c *Client) DownloadObject(bucket string, filename string, downloadFilename string) error {
	bs, err := c.GetObject(bucket, filename)
	if err != nil {
		return err
	}
	return os.WriteFile(downloadFilename, bs, 0644)
}

func (c *Client) GetObject(bucket string, filename string) ([]byte, error) {
	buffer := manager.NewWriteAtBuffer([]byte{})

	_, err := manager.NewDownloader(c.S3Client).Download(context.TODO(), buffer, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
	})
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (c *Client) DeleteObject(bucket string, filename string) (*s3.DeleteObjectOutput, error) {
	return c.S3Client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{Bucket: aws.String(bucket), Key: aws.String(filename)})
}

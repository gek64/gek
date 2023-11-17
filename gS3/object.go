package gS3

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
)

func (s *Session) Upload(bucket string, filename string, data []byte) (result *s3manager.UploadOutput, err error) {
	uploader := s3manager.NewUploader(s.Session)
	return uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
		Body:   bytes.NewReader(data),
	})
}

func (s *Session) Download(bucket string, filename string, downloadFilename string) (n int64, err error) {
	downloader := s3manager.NewDownloader(s.Session)

	f, err := os.Create(downloadFilename)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	return downloader.Download(f, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
	})
}

func (s *Session) Read(bucket string, filename string) (result *s3.GetObjectOutput, err error) {
	client := s.NewS3Client()
	return client.GetObjectWithContext(aws.BackgroundContext(), &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
	})
}

func (s *Session) Delete(bucket string, filename string) (result *s3.DeleteObjectOutput, err error) {
	client := s.NewS3Client()
	return client.DeleteObjectWithContext(aws.BackgroundContext(), &s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
	})
}

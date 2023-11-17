package gS3

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
)

func (s *Session) UploadObject(bucket string, filename string, data []byte) (result *s3manager.UploadOutput, err error) {
	uploader := s3manager.NewUploader(s.Session)
	return uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
		Body:   bytes.NewReader(data),
	})
}

func (s *Session) DownloadObject(bucket string, filename string, downloadFilename string) (n int64, err error) {
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

func (s *Session) PutObject(bucket string, filename string, data []byte) (result *s3.PutObjectOutput, err error) {
	client := s.NewS3Client()
	return client.PutObjectWithContext(aws.BackgroundContext(), &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
		Body:   aws.ReadSeekCloser(bytes.NewReader(data)),
	})
}

func (s *Session) GetObject(bucket string, filename string) (result *s3.GetObjectOutput, err error) {
	client := s.NewS3Client()
	return client.GetObjectWithContext(aws.BackgroundContext(), &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
	})
}

func (s *Session) DeleteObject(bucket string, filename string) (result *s3.DeleteObjectOutput, err error) {
	client := s.NewS3Client()
	return client.DeleteObjectWithContext(aws.BackgroundContext(), &s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
	})
}

func (s *Session) HeadObject(bucket string, filename string) (result *s3.HeadObjectOutput, err error) {
	client := s.NewS3Client()
	return client.HeadObjectWithContext(aws.BackgroundContext(), &s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
	})
}

func (s *Session) ListObjectsV2(bucket string) (result *s3.ListObjectsV2Output, err error) {
	client := s.NewS3Client()
	return client.ListObjectsV2WithContext(aws.BackgroundContext(), &s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
	})
}

func (s *Session) ListObjectVersions(bucket string, filename string) (result *s3.ListObjectVersionsOutput, err error) {
	client := s.NewS3Client()
	return client.ListObjectVersionsWithContext(aws.BackgroundContext(), &s3.ListObjectVersionsInput{
		Bucket: aws.String(bucket),
		Prefix: aws.String(filename),
	})
}

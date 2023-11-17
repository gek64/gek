package gS3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func (s *Session) CreateBucket(bucket string) (result *s3.CreateBucketOutput, err error) {
	client := s.NewS3Client()
	return client.CreateBucketWithContext(aws.BackgroundContext(), &s3.CreateBucketInput{
		Bucket: aws.String(bucket),
	})
}

func (s *Session) DeleteBucket(bucket string) (result *s3.DeleteBucketOutput, err error) {
	client := s.NewS3Client()
	return client.DeleteBucketWithContext(aws.BackgroundContext(), &s3.DeleteBucketInput{
		Bucket: aws.String(bucket),
	})
}

func (s *Session) HeadBucket(bucket string) (result *s3.HeadBucketOutput, err error) {
	client := s.NewS3Client()
	return client.HeadBucketWithContext(aws.BackgroundContext(), &s3.HeadBucketInput{
		Bucket: aws.String(bucket),
	})
}

func (s *Session) ListBuckets() (result *s3.ListBucketsOutput, err error) {
	client := s.NewS3Client()
	return client.ListBucketsWithContext(aws.BackgroundContext(), &s3.ListBucketsInput{})
}

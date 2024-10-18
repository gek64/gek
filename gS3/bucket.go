package gS3

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (c *Client) CreateBucket(bucket string) (*s3.CreateBucketOutput, error) {
	return c.S3Client.CreateBucket(context.TODO(), &s3.CreateBucketInput{Bucket: aws.String(bucket)})
}

func (c *Client) DeleteBucket(bucket string) (*s3.DeleteBucketOutput, error) {
	return c.S3Client.DeleteBucket(context.TODO(), &s3.DeleteBucketInput{Bucket: aws.String(bucket)})
}

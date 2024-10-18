package gS3

import (
	"crypto/tls"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"net/http"
)

type Client struct {
	S3Client *s3.Client
}

func NewS3Client(endpoint string, region string, accessKeyId string, secretAccessKey string, stsToken string, usePathStyle bool, allowInsecure bool) (client *Client) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: allowInsecure},
	}
	httpClient := http.Client{Transport: tr}

	return &Client{
		S3Client: s3.New(s3.Options{
			BaseEndpoint: aws.String(endpoint),
			Region:       region,
			Credentials:  credentials.NewStaticCredentialsProvider(accessKeyId, secretAccessKey, stsToken),
			HTTPClient:   &httpClient,
		}, func(options *s3.Options) {
			options.UsePathStyle = usePathStyle
		}),
	}
}

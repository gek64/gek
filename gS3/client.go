package gS3

import (
	"crypto/tls"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"net/http"
)

type Session struct {
	*session.Session
}

func NewS3Session(endpoint string, region string, accessKeyId string, secretAccessKey string, stsToken string, pathStyle bool, allowInsecure bool) (sess *Session) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: allowInsecure},
	}
	httpClient := http.Client{Transport: tr}

	return &Session{
		Session: session.Must(session.NewSession(&aws.Config{
			Endpoint:         aws.String(endpoint),
			Region:           aws.String(region),
			Credentials:      credentials.NewStaticCredentials(accessKeyId, secretAccessKey, stsToken),
			S3ForcePathStyle: aws.Bool(pathStyle),
			HTTPClient:       &httpClient,
		})),
	}
}

func (s *Session) NewS3Client() (client *s3.S3) {
	return s3.New(s.Session)
}

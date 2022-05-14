package infrastructure

import (
	"encoding/json"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// S3Config is a configuration for S3 management.
type S3Config struct {
	ID     string `json:"aws.s3.id"`
	Secret string `json:"aws.s3.secret"`
	Region string `json:"aws.s3.region"`
}

func ConnectToAws() *s3.S3 {
	secretManagerJSON := os.Getenv("APP_CONFIG_JSON")
	s3Config := S3Config{}
	json.Unmarshal([]byte(secretManagerJSON), &s3Config)

	creds := credentials.NewStaticCredentials(s3Config.ID, s3Config.Secret, "")
	_, err := creds.Get()
	if err != nil {
		panic(err)
	}

	cfg := aws.NewConfig().WithRegion(s3Config.Region).WithCredentials(creds)
	s3Connection := s3.New(session.New(), cfg)
	return s3Connection
}

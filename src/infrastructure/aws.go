package infrastructure

import (
	"encoding/json"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

// S3Config is a configuration for S3 management.
type S3Config struct {
	ID     string `json:"aws.s3.id"`
	Secret string `json:"aws.s3.secret"`
	Region string `json:"aws.s3.region"`
}

func ConnectToAws() *session.Session {
	secretManagerJSON := os.Getenv("APP_CONFIG_JSON")
	s3Config := S3Config{}
	json.Unmarshal([]byte(secretManagerJSON), &s3Config)

	awsSession, err := session.NewSession(
		&aws.Config{
			Region: aws.String(s3Config.Region),
			Credentials: credentials.NewStaticCredentials(
				s3Config.ID,
				s3Config.Secret,
				"",
			),
		})

	if err != nil {
		panic(err)
	}

	return awsSession
}

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

var AccessKeyID string
var SecretAccessKey string
var MyRegion string

func ConnectToAws() *session.Session {
	secretManagerJSON := os.Getenv("APP_CONFIG_JSON")
	s3Config := S3Config{}
	json.Unmarshal([]byte(secretManagerJSON), &s3Config)

	AccessKeyID = s3Config.ID
	SecretAccessKey = s3Config.Secret
	MyRegion = s3Config.Region
	awsSession, err := session.NewSession(
		&aws.Config{
			Region: aws.String(MyRegion),
			Credentials: credentials.NewStaticCredentials(
				AccessKeyID,
				SecretAccessKey,
				"",
			),
		})

	if err != nil {
		panic(err)
	}

	return awsSession
}

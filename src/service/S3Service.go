package service

import (
	"bytes"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/yumaeda/sakaba-api/src/infrastructure"
)

// S3Service handle AWS S3 operations.
type S3Service struct{}

// Upload uploads the file to AWS S3.
func (c *S3Service) Upload(restaurantID string, fileName string, file []byte) error {
	s3Bucket := "admin.tokyo-takeout.com"
	awsSession := infrastructure.ConnectToAws()
	objectKey := fmt.Sprintf("images/restaurants/%s/%s.jpeg", restaurantID, fileName)
	uploadParams := &s3.PutObjectInput{
		Bucket: aws.String(s3Bucket),
		ACL:    aws.String("public-read"),
		Key:    aws.String(objectKey),
		Body:   bytes.NewReader(file),
	}

	_, err := awsSession.PutObject(uploadParams)
	return err
}

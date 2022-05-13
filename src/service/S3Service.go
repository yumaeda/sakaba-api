package service

import (
	"bytes"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"sakaba.link/api/src/infrastructure"
)

// S3Service handle AWS S3 operations.
type S3Service struct{}

// Upload uploads the file to AWS S3.
func (c *S3Service) Upload(restaurantID string, fileName string, file []byte) (*s3manager.UploadOutput, error) {
	s3Bucket := "admin.tokyo-takeout.com"
	awsSession := infrastructure.ConnectToAws()
	filePath := fmt.Sprintf("images/restaurants/%s/%s.jpeg", restaurantID, fileName)
	uploader := s3manager.NewUploader(awsSession)

	return uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(s3Bucket),
		ACL:    aws.String("public-read"),
		Key:    aws.String(filePath),
		Body:   bytes.NewReader(file),
	})
}

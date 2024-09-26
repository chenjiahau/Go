package util

import (
	"context"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// Create a new session with AWS S3
func CreateSession(region, accessKeyId, secretAccessKey, bucketName string) (*session.Session, error) {
	return session.NewSession(&aws.Config{
  	Region: aws.String(region),
  	Credentials: credentials.NewStaticCredentials(
			accessKeyId,
   		secretAccessKey,
   		"",
  	),
	})
}

// Upload file to AWS S3
func UploadFileToS3(file multipart.File, fileType, filePath, fileName, fileExtension, region, accessKeyId, secretAccessKey, bucketName string) (*s3manager.UploadOutput, error) {
	s3Session, err := CreateSession(region, accessKeyId, secretAccessKey, bucketName)
	if err != nil {
		return nil, err
	}

	uploader := s3manager.NewUploader(s3Session)
	input := &s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key: aws.String(filePath + "/" + fileName + fileExtension),
		Body: file,
		ContentType: aws.String(fileType),
	}

	output, err := uploader.UploadWithContext(context.Background(), input)
	if err != nil {
		return nil, err
	}

	return output, nil
}
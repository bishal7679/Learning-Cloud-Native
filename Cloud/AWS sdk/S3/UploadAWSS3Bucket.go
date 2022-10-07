package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func UploadFile(uploader *s3manager.Uploader, filePath string, fileName string, bucketName string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileName),
		Body:   file,
	})

	return err
}
func main() {
	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: "default",
		Config: aws.Config{
			Region: aws.String("us-east-1"),
		},
	})

	if err != nil {
		fmt.Printf("Failed to initialize new session %v", err)
		return
	}

	bucketName := "learning-go-sdk-s3"
	uploader := s3manager.NewUploader(sess)
	fileName := "pic.png"

	err = UploadFile(uploader, "pic.png", fileName, bucketName)
	if err != nil {
		fmt.Printf("Error occured uploading the file %v", err)
	}

	fmt.Println("Successfully uploaded the file!")

}

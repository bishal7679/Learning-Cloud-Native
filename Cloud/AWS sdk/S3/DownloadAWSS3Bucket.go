package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func DownloadFile(downloader *s3manager.Downloader, bucketName string, key string) error {
	file, err := os.Create(key)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = downloader.Download(file, &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
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
	downloader := s3manager.NewDownloader(sess)
	key := "pic.png"

	err = DownloadFile(downloader, bucketName, key)
	if err != nil {
		fmt.Printf("Error occured downloading the file %v", err)
	}

	fmt.Println("Successfully Downloaded the file!")

}

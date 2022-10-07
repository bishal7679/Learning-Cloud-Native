package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func ListBucket(client *s3.S3) (*s3.ListBucketsOutput, error) {
	res, err := client.ListBuckets(nil)
	if err != nil {
		return nil, err
	}

	return res, nil

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

	s3Client := s3.New(sess)

	buckets, err := ListBucket(s3Client)

	if err != nil {
		fmt.Printf("Couldn't List buckets %v", err)
	}

	for _, bucket := range buckets.Buckets {
		fmt.Printf("found bucket: %s  , Created at: %s\n", *bucket.Name, *bucket.CreationDate)
	}
}

package client

// package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Client interface {
	CopyObject(sourceBucketName, objectKyey, destBucketName, targetKeyArn string) error
}

type s3Client struct {
	bucketClient *s3.Client
}

func NewS3Client() S3Client {
	region := os.Getenv("AWS_REGION")
	config, _ := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	return &s3Client{bucketClient: s3.NewFromConfig(config)}
}

func (client *s3Client) CopyObject(sourceBucketName, objectKyey, destBucketName, targetKeyArn string) error {
	_, err := client.bucketClient.CopyObject(context.TODO(), &s3.CopyObjectInput{
		Bucket:     aws.String(destBucketName),
		CopySource: aws.String(fmt.Sprintf("%s/%s", sourceBucketName, objectKyey)),
		Key:        aws.String(objectKyey),
	})
	return err
}

// func main() {
// 	os.Setenv("AWS_REGION", "us-east-1")
// 	s3Client := NewS3Client()
// 	sourceBucketName := "sample-asd-123"
// 	objectKey := "sample.txt"
// 	destBucketName := "sample-asd-234"
// 	targetKeyArn := ""
// 	err := s3Client.CopyObject(sourceBucketName, objectKey, destBucketName, targetKeyArn)
// 	if err != nil {
// 		fmt.Println("Error occurred: ", err)
// 	} else {
// 		fmt.Println("Copy done!")
// 	}
// }

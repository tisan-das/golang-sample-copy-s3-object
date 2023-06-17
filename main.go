package main

import (
	"copy-s3-object/client"
	"copy-s3-object/router"
	"copy-s3-object/service"
	"net/http"
	"os"
)

func main() {
	os.Setenv("AWS_REGION", "us-east-1")

	s3Client := client.NewS3Client()
	s3Service := service.NewS3Service(s3Client)
	router := router.NewWebRouter(s3Service)
	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}

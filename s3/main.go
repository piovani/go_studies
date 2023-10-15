package main

import (
	"fmt"
	"io"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func main() {
	file, _ := os.Open("./test.json")

	Upload(file)
}

func Upload(file io.Reader) {

	sess := session.Must(session.NewSession(
		&aws.Config{
			Credentials:      credentials.NewStaticCredentials("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY", ""),
			Region:           aws.String("us-east-1"),
			S3ForcePathStyle: aws.Bool(true),
		},
	))

	// S3 service client the Upload manager will use.
	s3Svc := s3.New(sess)

	// Create an uploader with S3 client and default options
	uploader := s3manager.NewUploaderWithClient(s3Svc)

	upParams := &s3manager.UploadInput{
		Bucket: aws.String("mybucket"),
		Key:    aws.String("teste2.json"),
		Body:   file,
	}

	// Perform an upload.
	result, err := uploader.Upload(upParams)

	fmt.Println(result, err)

}

package main

import (
	"os"
	"log"
	"fmt"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Profile: "sandbox",
		SharedConfigState: session.SharedConfigEnable,
	}))

	bucket := "rf-sandbox-ansiblelog"
	filename := "./foo"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	
	defer file.Close()
	
	uploader := s3manager.NewUploader(sess)

	result, err := uploader.Upload(&s3manager.UploadInput{
	    Bucket: aws.String(bucket),
	    Key: aws.String(filename),
	    Body: file,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("file uploaded to, %s\n", aws.StringValue(&result.Location))
}

package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/jetoneza/personal_website/pkg/config"
)

func main() {
	timeout := time.Duration(time.Minute * 2)

	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(config.AWS_REGION),
	}))
	svc := s3.New(sess)

	// Create a context with a timeout that will abort the upload if it takes
	// more than the passed in timeout.
	ctx := context.Background()
	var cancelFn func()
	if timeout > 0 {
		ctx, cancelFn = context.WithTimeout(ctx, timeout)
	}

	// Ensure the context is canceled to prevent leaking.
	if cancelFn != nil {
		defer cancelFn()
	}

	_, err := svc.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Bucket: aws.String(config.S3_BUCKET),
		Key:    aws.String(config.S3_OBJECT_KEY),
		Body:   os.Stdin,
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok && aerr.Code() == request.CanceledErrorCode {
			fmt.Fprintf(os.Stderr, "AWS: upload canceled due to timeout, %v\n", err)
		} else {
			fmt.Fprintf(os.Stderr, "AWS: failed to upload object, %v\n", err)
		}
		os.Exit(1)
	}

	fmt.Printf("AWS: successfully uploaded file\n")
}

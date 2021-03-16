package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"cloud.google.com/go/bigquery"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"google.golang.org/api/option"
)

//client initialized globally to be accessed by other functions
var client *bigquery.Client
var ctx context.Context

//Lambda accepts init functions and runs them like a regular Go program https://docs.aws.amazon.com/lfilesambda/latest/dg/lambda-golang.html
func init() {
	svc := s3.New(session.New())
	input := &s3.GetObjectInput{
		Bucket: aws.String("jsonfiles312021"),
		Key:    aws.String("Project-8a8c500b8c6d.json"),
	}
	result, _ := svc.GetObject(input)
	defer result.Body.Close()
	body, _ := ioutil.ReadAll(result.Body)

	ctx = context.Background()

	var err error
	client, err = bigquery.NewClient(ctx, "first-vision-305321", option.WithCredentialsJSON(body))
	if err != nil {
		log.Fatalf("bigquery.NewClient: %v", err)
	}

}

func main() {
	lambda.Start(handler)
	fmt.Printf("hello")
}

func handler() {

}

package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var db = dynamodb.New(session.New(), aws.NewConfig().WithRegion("us-east-1"))

func getPosts() (int64, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String("margarine-posts"),
	}

	result, err := db.Scan(input)
	if err != nil {
		return 0, err
	}

	return *result.Count, nil
}

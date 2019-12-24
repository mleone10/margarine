package posts

import (
	"log"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Post struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
}

var db = dynamodb.New(session.New(), aws.NewConfig().WithRegion("us-east-1"))

func GetPosts() ([]Post, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String("margarine-posts"),
	}

	result, err := db.Scan(input)
	if err != nil {
		return nil, err
	}

	posts := []Post{}

	for _, i := range result.Items {
		post := Post{}
		err = dynamodbattribute.UnmarshalMap(i, &post)

		if err != nil {
			log.Println(err.Error())
		} else {
			posts = append(posts, post)
		}
	}

	return posts, nil
}

func GetPost(id int) (*Post, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String("margarine-posts"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				N: aws.String(strconv.Itoa(id)),
			},
		},
	}

	result, err := db.GetItem(input)
	if err != nil {
		return nil, err
	}
	if result.Item == nil {
		return nil, nil
	}

	post := Post{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &post)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

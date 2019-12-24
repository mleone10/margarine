package main

import (
	"encoding/json"
	"github.com/mleone10/margarine/internal/posts"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response events.APIGatewayProxyResponse
type Request events.APIGatewayProxyRequest

func Handler(req Request) (Response, error) {
	posts, err := posts.GetPosts()
	if err != nil {
		body, _ := json.Marshal(map[string]interface{}{
			"message": "Failed to retrieve posts from database",
		})
		return Response{
			StatusCode: 500,
			Body:       string(body),
		}, nil
	}

	body, err := json.Marshal(map[string]interface{}{
		"posts": posts,
	})

	if err != nil {
		return Response{StatusCode: 500}, nil
	}

	resp := Response{
		StatusCode: 200,
		Body:       string(body),
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}

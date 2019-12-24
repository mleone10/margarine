package main

import (
	"encoding/json"
	"github.com/mleone10/margarine/internal/posts"
	"github.com/mleone10/margarine/internal/response"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Request events.APIGatewayProxyRequest

func Handler(req Request) (response.Response, error) {
	posts, err := posts.GetPosts()
	if err != nil {
		return response.ServerErrorWithMessage("Failed to retrieve posts from database"), nil
	}

	body, err := json.Marshal(map[string]interface{}{
		"posts": posts,
	})

	if err != nil {
		return response.ServerError(), nil
	}

	return response.Response{
		StatusCode: 200,
		Body:       string(body),
	}, nil
}

func main() {
	lambda.Start(Handler)
}

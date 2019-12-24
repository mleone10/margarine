package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Post struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
}

type Response events.APIGatewayProxyResponse
type Request events.APIGatewayProxyRequest

func Handler(req Request) (Response, error) {
	posts, err := getPosts()
	if err != nil {
		return Response{
			StatusCode: 500,
			Body: safeMarshal(map[string]interface{}{
				"message": "Failed to retrieve posts from database",
			}),
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

func safeMarshal(j map[string]interface{}) string {
	s, _ := json.Marshal(j)
	return string(s)
}

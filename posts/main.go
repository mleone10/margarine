package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response events.APIGatewayProxyResponse
type Request events.APIGatewayProxyRequest

func Handler(req Request) (Response, error) {
	count, _ := getPosts()

	body, err := json.Marshal(map[string]interface{}{
		"count": count,
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

package main

import (
	"encoding/json"
	"github.com/mleone10/margarine/internal/posts"
	"github.com/mleone10/margarine/internal/response"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Request events.APIGatewayProxyRequest

func Handler(req Request) (response.Response, error) {
	idParam := req.PathParameters["id"]

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return response.ClientErrorWithMessage("ID param must be a valid integer"), nil
	}

	post, err := posts.GetPost(id)
	if err != nil {
		return response.ServerErrorWithMessage("Failed to retrieve post from database"), nil
	}
	if post == nil {
		return response.NotFound(), nil
	}

	body, err := json.Marshal(post)

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

package response

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

type Response events.APIGatewayProxyResponse

func ServerError() Response {
	return Response{
		StatusCode: 500,
	}
}

func ServerErrorWithMessage(message string) Response {
	body, _ := json.Marshal(map[string]interface{}{
		"message": message,
	})
	return Response{
		StatusCode: 500,
		Body:       string(body),
	}
}

func ClientErrorWithMessage(message string) Response {
	body, _ := json.Marshal(map[string]interface{}{
		"message": message,
	})
	return Response{
		StatusCode: 400,
		Body:       string(body),
	}
}

func NotFound() Response {
	return Response{
		StatusCode: 404,
	}
}

func Success(body string) Response {
	return Response{
		StatusCode: 200,
		Body:       body,
		Headers: map[string]string{
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Credentials": "true",
		},
	}
}

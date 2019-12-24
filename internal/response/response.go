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

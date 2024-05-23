package main

import (
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var pathHandlers = map[string]func(map[string]interface{}) (string, string, error){
	"uala": uala,
}

func getPathHead(path string) string {
	// Remove leading '/'
	trimmedPath := strings.TrimPrefix(path, "/")
	// Find the position of the second '/' character
	secondSlashIndex := strings.Index(trimmedPath, "/")

	// Extract the first path component
	if secondSlashIndex != -1 {
		return trimmedPath[:secondSlashIndex]
	}
	return trimmedPath
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	pathHead := getPathHead(request.Path)
	requestData := make(map[string]interface{})

	handlerFunc, exists := pathHandlers[pathHead]
	if !exists {
		handlerFunc = fallback
	}

	contentType, body, err := handlerFunc(requestData)
	var statusCode int
	if err == nil {
		statusCode = 200
	} else {
		statusCode = 500
	}

	return events.APIGatewayProxyResponse{
		Body:       body,
		StatusCode: statusCode,
		Headers: map[string]string{
			"Content-Type": contentType,
		},
	}, nil
}

func main() {
	lambda.Start(handler)
}

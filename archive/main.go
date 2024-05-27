package kokuban

import (
	"fmt"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type AppResponse struct {
	Body        string
	ContentType string
}

type AppRequest struct {
}

type Path = string

var pathHandlers = map[Path]func(request AppRequest) (AppResponse, error){
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

	handlerFunc, exists := pathHandlers[pathHead]
	if !exists {
		handlerFunc = fallback
	}

	req := AppRequest{}

	resp, err := handlerFunc(req)
	var statusCode int
	if err == nil {
		statusCode = 200
	} else {
		statusCode = 500
	}

	fmt.Println("Returning RESPNSE: ....")
	return events.APIGatewayProxyResponse{
		Body:       resp.Body,
		StatusCode: statusCode,
		Headers: map[string]string{
			"Content-Type": resp.ContentType,
		},
	}, nil
}

func main() {
	lambda.Start(handler)
}

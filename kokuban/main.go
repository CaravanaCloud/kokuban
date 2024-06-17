package main

import (
	"strings"

	"net/http"

	"github.com/gin-gonic/gin"
)

type AppResponse struct {
	Body        string
	ContentType string
}

type AppRequest struct {
}

type Path = string

// TODO: Map path handlers
// var pathHandlers = map[Path]func(request AppRequest) (AppResponse, error){
//	"uala": uala,
// }

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

func main() {
	r := gin.Default()
	r.NoRoute( func(c *gin.Context) {
		// TODO: Design activities API / Frontend / ... (create those tickets)
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

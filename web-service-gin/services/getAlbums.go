package services

import (
	"net/http"
	"go-guide/web-service-gin/types"
	"github.com/gin-gonic/gin"
)
// GetAlbums function that takes a gin.Context parameter. 
// gin.Context is the most important part of Gin. 
// It carries request details, validates and serializes JSON, and more. 
// (Despite the similar name, this is different from Go’s built-in context package.)
// GetAlbums responds with the list of all albums as JSON.
func GetAlbums(c *gin.Context) {
   	// Call Context.IndentedJSON to serialize the struct into JSON and add it to the response.
	// The function’s first argument is the HTTP status code you want to send to the client. 
	// Here, you’re passing the StatusOK constant from the net/http package to indicate 200 OK.
	c.IndentedJSON(http.StatusOK, types.Albums)
}
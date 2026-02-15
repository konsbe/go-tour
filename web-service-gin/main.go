package main

import (
	"go-guide/web-service-gin/services"
    "github.com/gin-gonic/gin"
)

func main() {
	// Initialize a Gin router using Default.
    router := gin.Default()
	// Use the GET function to associate the GET HTTP method and /albums path with a handler function.
    router.GET("/albums", services.GetAlbums)
	// Associate the POST method at the /albums path with the postAlbums function.
	// With Gin, you can associate a handler with an HTTP method-and-path combination. 
	// In this way, you can separately route requests sent to a single path based on the method the client is using.
	router.POST("/albums", services.PostAlbums)
	// Associate the GET method at the /albums/:id path with the getAlbumByID function.
	// The :id syntax in the path indicates that this is a path parameter, 
	// and that the part of the URL that matches this parameter will be passed to the handler function.
	router.GET("/albums/:id", services.GetAlbumByID)

	// Use the Run function to attach the router to an http.Server and start the server.
    router.Run("localhost:8080")
}
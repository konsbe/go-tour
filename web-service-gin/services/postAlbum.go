package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-guide/web-service-gin/types"
)

// PostAlbums adds an album from JSON received in the request body.
func PostAlbums(c *gin.Context) {
    var newAlbum types.Album

    // Call BindJSON to bind the received JSON to
    // newAlbum.
    if err := c.BindJSON(&newAlbum); err != nil {
		// If the JSON is not valid, return a 400 error with a message.
        c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
        return
    }

	// Append the album struct initialized from the JSON to the albums slice.
	types.Albums = append(types.Albums, newAlbum)
	// Add a 201 status code to the response, along with JSON representing the album you added.
    c.IndentedJSON(http.StatusCreated, newAlbum)
}
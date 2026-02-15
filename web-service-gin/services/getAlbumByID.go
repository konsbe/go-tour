package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-guide/web-service-gin/types"
)

// GetAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func GetAlbumByID(c *gin.Context) {
	// Use Context.Param to retrieve the id path parameter from the URL. 
	// When you map this handler to a path, you’ll include a placeholder for the parameter in the path.
    id := c.Param("id")

    
	// Loop over the album structs in the slice, 
	// looking for one whose ID field value matches the id parameter value. 
	// If it’s found, you serialize that album struct to JSON and return it as a response with a 200 OK HTTP code.
    for _, a := range types.Albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
	// Return an HTTP 404 error with http.StatusNotFound if the album isn’t found.
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
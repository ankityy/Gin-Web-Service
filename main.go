package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main(){

	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.POST("/albums", postAlbums)
	router.PATCH("/albums/:id", updateAlbumById)
	router.DELETE("/albums/:id", deleteAlbumById)

	router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context){
	var newAlbum album

	// Call BindJSON to bind the received JSON to
    // newAlbum.
	err := c.BindJSON(&newAlbum)
	if err != nil {
		return
	}
	//append the newalbum in albums
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, albums)
}

func getAlbumById(c *gin.Context){
	id := c.Param("id")
	
	// Loop over the list of albums, looking for
    // an album whose ID value matches the parameter.

	for _, album := range albums {
		if id == album.ID{
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No album matching"})

}

// updateAlbumById 
func updateAlbumById(c *gin.Context){
	id := c.Param("id")

	//json update album
	var updateAlbum album 
	err := c.BindJSON(&updateAlbum)
	if err != nil {
		return
	}
	// Loop over all the albums and update the album object
	for key, album := range albums{
		if album.ID == id {
			albums[key] = updateAlbum
			c.IndentedJSON(http.StatusOK, albums)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No album matching"})
}

// deleteAlbumById
func deleteAlbumById(c *gin.Context){
	id := c.Param("id")

	for key, album := range albums{
		if album.ID == id {

			albums = append(albums[:key],albums[key+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Album Deleted"})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No album matching"})

}


package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.GET("/albums/getMostExpensiveAlbum", getMostExpensiveAlbum)
	router.GET("/albums/getCheapestAlbum", getCheapestAlbum)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}

// returns the slice of all the albums to the user
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// adds an album to the slice of albums.
func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbumById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func getMostExpensiveAlbum(c *gin.Context) {
	max := albums[0].Price
	maxAlbum := albums[0]

	for _, a := range albums {
		if a.Price > max {
			max = a.Price
			maxAlbum = a
		}
	}

	c.IndentedJSON(http.StatusOK, maxAlbum)
}

func getCheapestAlbum(c *gin.Context) {
	min := albums[0].Price
	minAlbum := albums[0]

	for _, a := range albums {
		if a.Price < min {
			min = a.Price
			minAlbum = a
		}
	}

	c.IndentedJSON(http.StatusOK, minAlbum)
}

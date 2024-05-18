package main

import (
	"github.com/gin-gonic/gin"
)

type voting struct {
	voted string `json:"voted"`
	voter string `json:"voter"`
}

var votings = []voting{
	{voted: "Axl"},
	{voted: "Jil"},
	{voted: "Kim"},
	{voted: "Linora"},
	{voted: "Lulu"},
	{voted: "Lutzispatzi"},
	{voted: "Manuel"},
	{voted: "Patrick"},
}

func main() {
	router := gin.Default()
	router.GET("/result", getResult)
	router.POST("/vote", postVote)
	router.POST("/reset", postReset)

	router.Run("localhost:8080")
}

func appendVote(c *gin.Context) {

}

/*
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
*/

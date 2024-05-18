package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type voting struct {
	voted string `json:"voted"`
	voter string `json:"voter"`
}

var finishedVoting = false

var votings = []voting{}

func main() {
	router := gin.Default()
	//router.GET("/result", getResult)
	router.GET("/hello", hello_world)
	router.POST("/vote", appendVote)
	router.POST("/reset", reset)

	router.Run("localhost:8080")
}

func hello_world(c *gin.Context) {
	c.Status(http.StatusOK)
}

func reset(c *gin.Context) {
	votings = []voting{}
	c.Status(http.StatusOK)
}

func appendVote(c *gin.Context) {
	vter := c.Query("voter")
	vted := c.Query("voted")

	present := isVotedPresent(vted)
	if present > 0 {
		print("Present: ", present)
		v := votings[present-1].voter
		v = v + ", " + vter

		vot := voting{
			voted: vted,
			voter: v,
		}
		println("remove")
		remove(present - 1)
		println("append")
		votings = append(votings, vot)

	} else {
		vot := voting{
			voted: vted,
			voter: vter,
		}
		votings = append(votings, vot)
	}
	checkFinish()
	for _, v := range votings {
		println(v.voted+": ", v.voter)
	}

	c.Status(http.StatusOK)
}

func isVotedPresent(votedKey string) int {
	for i, v := range votings {
		i++
		println("index: ", i)
		if v.voted == votedKey {
			return i
		}
	}
	return -1
}

func remove(s int) {
	v := append(votings[:s], votings[s+1:]...)
	votings = v
}

func checkFinish() {
	count := 0
	for _, v := range votings {
		count += countWords(v.voter, ", ")
	}
	if count == 8 {
		println("Finished")
		finishedVoting = true
	}
}

func countWords(s string, sep string) int {
	words := strings.Split(s, sep)
	return len(words)
}

// Query string parameters are parsed using the existing underlying request object.
//   /vote?voted=Axl&voter=Jil

/*
	{voted: "Axl", voter: "dfdf"},
	{voted: "Jil"},
	{voted: "Kim"},
	{voted: "Linora"},
	{voted: "Lulu"},
	{voted: "Lutzispatzi"},
	{voted: "Manuel"},
	{voted: "Patrick"},




*/

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

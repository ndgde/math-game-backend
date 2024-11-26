package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ndgde/math-game-backend/cmd/db"
)

func main() {
	_ = db.NewDB()

	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8089")
}

func getAlbums(c *gin.Context) {
	albums := db.GetAllAlbums()
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum db.Album
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "the request does not match the json format"})
		return
	}

	if err := db.CreateAlbum(newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("The request format is not expected %s", err)})
	}

	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	strId := c.Param("id")

	id64, err := strconv.ParseInt(strId, 10, 32)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	if id64 < -2147483648 || id64 > 2147483647 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Number is out of range int32"})
		return
	}
	id := int32(id64)

	albums := db.GetAlbumByID(id)

	if len(albums) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not fount"})
		return
	}

	c.IndentedJSON(http.StatusOK, albums)
}

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}

var albums = []album {
	{ID: "1", Title: "Yoru ni kakeru", Artist: "Yoasobi", Price: 10.99},
	{ID: "2", Title: "Kartoyono medot janjimu", Artist: "Cak Nan", Price: 11.99},
	{ID: "3", Title: "Sparkle", Artist: "Radwimps", Price: 9.99},
}

func getAlbums(c *gin.Context)  {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context)  {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumById(c *gin.Context)  {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message":"album not found"})
}

func removeAlbumById(c *gin.Context)  {
	id := c.Param("id")

	for i, a := range albums {
		if a.ID == id {
			// remove
			albums = append(albums[:i], albums[i+1:]...)
			c.IndentedJSON(http.StatusOK, albums)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message":"album not found"})
}

func main()  {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.DELETE("/albums/:id", removeAlbumById)
	router.Run("localhost:3000")
}

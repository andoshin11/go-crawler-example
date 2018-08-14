package main

import (
	"net/http"

	"github.com/andoshin11/go-crawler-example/src/handler"
	"github.com/gin-gonic/gin"
)

func test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"test": "hoge"})
}

func main() {
	router := gin.Default()

	router.GET("/test", test)

	// Crawler namespace
	clawl := router.Group("/crawl")
	{
		clawl.GET("/items", handler.CrawlArtscapeItems)
		clawl.GET("/item", handler.CrawlArtscapeItem)
	}

	router.Run(":8080")
}

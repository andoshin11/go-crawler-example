package handler

import (
	"context"
	"log"

	"github.com/andoshin11/go-crawler-example/src/client"
	"github.com/andoshin11/go-crawler-example/src/domain/worker"
	"github.com/gin-gonic/gin"
)

// CrawlArtscapeItems is the request handler
func CrawlArtscapeItems(c *gin.Context) {
	ctx := context.Background()
	client, err := client.NewFirestoreClient(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	worker := worker.NewArtscapeWorker(client)
	worker.CrawlItems(ctx)
	return
}

// CrawlArtscapeItem is the request handler
func CrawlArtscapeItem(c *gin.Context) {
	ctx := context.Background()
	client, err := client.NewFirestoreClient(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	worker := worker.NewArtscapeWorker(client)
	worker.CrawlDetail(ctx)
	return
}

// // CrawlExhibitionListHandler is the request handler
// func CrawlExhibitionListHandler(ctx context.Context) {
// 	client, err := client.NewFirestoreClient(ctx)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	worker := worker.NewArtscapeWorker(client)
// 	worker.CrawlExhibitionList(ctx)
// 	return
// }

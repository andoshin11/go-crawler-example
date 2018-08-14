package worker

import (
	"context"
	"log"
	"strconv"

	"cloud.google.com/go/firestore"
	"github.com/andoshin11/go-crawler-example/src/domain/fetcher"
	"github.com/andoshin11/go-crawler-example/src/domain/parser"
	"github.com/andoshin11/go-crawler-example/src/domain/uploader"
	"github.com/andoshin11/go-crawler-example/src/repository"
	"github.com/andoshin11/go-crawler-example/src/types"
)

// ArtscapeWorker interface
type ArtscapeWorker interface {
	CrawlItems(ctx context.Context)
	CrawlDetail(ctx context.Context)
}

type artscapeWorker struct {
	Client *firestore.Client
}

// NewArtscapeWorker returns struct
func NewArtscapeWorker(Client *firestore.Client) ArtscapeWorker {
	return &artscapeWorker{Client}
}

func (w *artscapeWorker) CrawlItems(ctx context.Context) {
	parser := parser.NewArtscapeParser()
	fetcher := fetcher.NewArtscapeFetcher(parser)
	museumRepository := repository.NewMuseumRepository(w.Client)
	uploader := uploader.NewArtscapeUploader(museumRepository)

	// worker count
	fetcherWc := 0
	uploaderWc := 0

	chs := types.NewChannels()

	// 47都道府県の各エリア
	for i := 1; i <= 4; i++ {
		fetcherWc++
		id := strconv.Itoa(i)
		url := "http://artscape.jp/mdb/mdb_result.php?area=" + id
		go fetcher.FetchItems(url, chs)
	}

	done := false
	for !done {
		select {
		case res := <-chs.FetcherResult:
			link := res.URL
			subIdentifier := link[23 : len(link)-10] // parse id

			uploaderWc++
			go uploader.RegisterArtscapeMuseum(ctx, subIdentifier, chs)
		case <-chs.FetcherDone:
			fetcherWc--
			if fetcherWc == 0 && uploaderWc == 0 {
				done = true
			}
		case <-chs.UploaderDone:
			uploaderWc--
			if fetcherWc == 0 && uploaderWc == 0 {
				done = true
			}
		}
	}
}

func (w *artscapeWorker) CrawlDetail(ctx context.Context) {
	parser := parser.NewArtscapeParser()
	fetcher := fetcher.NewArtscapeFetcher(parser)
	museumRepository := repository.NewMuseumRepository(w.Client)
	uploader := uploader.NewArtscapeUploader(museumRepository)

	// worker count
	fetcherWc := 0
	uploaderWc := 0

	chs := types.NewDetailChannels()

	museums, err := museumRepository.GetAll(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	for _, museum := range museums {
		if museum != nil {
			fetcherWc++
			go fetcher.FetchDetail(ctx, museum.SubIdentifier, museum.Identifier, chs)
		}
	}

	done := false
	for !done {
		select {
		case res := <-chs.FetcherResult:
			uploaderWc++
			go uploader.UpdateArtscapeMuseum(ctx, res.ID, res.Item, chs)
		case <-chs.FetcherDone:
			fetcherWc--
			if fetcherWc == 0 && uploaderWc == 0 {
				done = true
			}
		case <-chs.UploaderDone:
			uploaderWc--
			if fetcherWc == 0 && uploaderWc == 0 {
				done = true
			}
		}
	}
}

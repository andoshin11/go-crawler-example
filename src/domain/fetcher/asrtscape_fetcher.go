package fetcher

import (
	"context"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
	"github.com/andoshin11/go-crawler-example/src/domain/parser"
	"github.com/andoshin11/go-crawler-example/src/types"
)

// ArtscapeFetcher interface
type ArtscapeFetcher interface {
	FetchItems(u string, ch *types.Channels) (err error)
	FetchDetail(ctx context.Context, subIdentifier, identifier string, ch *types.DetailChannels) (err error)
}

type artscapeFetcher struct {
	parser parser.ArtscapeParser
}

// NewArtscapeFetcher returns struct
func NewArtscapeFetcher(parser parser.ArtscapeParser) ArtscapeFetcher {
	return &artscapeFetcher{parser}
}

// FetchItems returns the list of child page urls
func (f *artscapeFetcher) FetchItems(u string, ch *types.Channels) (err error) {
	defer func() { ch.FetcherDone <- 0 }()

	baseURL, err := url.Parse(u)
	if err != nil {
		return
	}

	resp, err := http.Get(baseURL.String())
	if err != nil {
		return
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return
	}

	urls := f.parser.ParseItems(doc)

	for _, item := range urls {
		itemURL, err := baseURL.Parse(item)
		if err == nil {
			ch.FetcherResult <- types.FetcherResult{
				URL: itemURL.String(),
			}
		}
	}
	return
}

// FetchDetail returns the item detail
func (f *artscapeFetcher) FetchDetail(ctx context.Context, subIdentifier, identifier string, ch *types.DetailChannels) (err error) {
	defer func() { ch.FetcherDone <- 0 }()
	path := "http://artscape.jp/mdb/" + subIdentifier + "_1900.html"
	resp, err := http.Get(path)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return
	}

	museum := f.parser.ParseDetail(doc)

	ch.FetcherResult <- types.DetailFetcherResult{
		ID:   identifier,
		Item: museum,
	}

	return
}

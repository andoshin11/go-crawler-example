package fetcher

import (
	"context"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
	"github.com/andoshin11/go-crawler-example/src/domain/parser"
	"github.com/andoshin11/go-crawler-example/src/types"
)

// ArtscapeItemsFetcher returns the list of child page urls
func ArtscapeItemsFetcher(u string, ch *types.Channels) (err error) {
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

	urls := parser.ArtscapeItemsParser(doc)

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

// ArtscapeItemFetcher returns the item detail
func ArtscapeItemFetcher(ctx context.Context, subIdentifier, identifier, parentID string, ch *types.DetailChannels) (err error) {
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

	museum := parser.ArtscapeItemParser(doc)

	ch.FetcherResult <- types.DetailFetcherResult{
		ID:       identifier,
		ParentID: parentID,
		Item:     museum,
	}

	return
}

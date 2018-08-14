package uploader

import (
	"context"
	"log"

	"github.com/andoshin11/go-crawler-example/src/repository"
	"github.com/andoshin11/go-crawler-example/src/types"
)

const collection = "crawlerExample"

// ArtscapeUploader interface
type ArtscapeUploader interface {
	RegisterArtscapeMuseum(ctx context.Context, subIdentifier string, ch *types.Channels)
	UpdateArtscapeMuseum(ctx context.Context, identifier string, museum *types.Museum, ch *types.DetailChannels)
}

type artscapeUploader struct {
	MuseumRepository repository.MuseumRepository
}

// NewArtscapeUploader returns struct
func NewArtscapeUploader(MuseumRepository repository.MuseumRepository) ArtscapeUploader {
	return &artscapeUploader{MuseumRepository}
}

func (u *artscapeUploader) RegisterArtscapeMuseum(ctx context.Context, subIdentifier string, ch *types.Channels) {
	_, err := u.MuseumRepository.GetBySubIdentifier(ctx, subIdentifier)
	if err != nil {
		log.Println("New Item: " + subIdentifier)
		err = u.MuseumRepository.AddItem(ctx, subIdentifier, "artscape")
		if err != nil {
			log.Fatalln(err)
		}
	}

	defer func() { ch.UploaderDone <- 0 }()
}

func (u *artscapeUploader) UpdateArtscapeMuseum(ctx context.Context, identifier string, museum *types.Museum, ch *types.DetailChannels) {
	err := u.MuseumRepository.UpdateItem(ctx, identifier, museum)
	if err != nil {
		log.Println(err)
	}

	defer func() { ch.UploaderDone <- 0 }()
}

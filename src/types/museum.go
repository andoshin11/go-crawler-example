package types

import (
	"time"
)

// Museum type
type Museum struct {
	Identifier    string    `firestore:"identifier"`
	CreatedAt     time.Time `firestore:"createdAt"`
	UpdatedAt     time.Time `firestore:"updatedAt"`
	Name          string    `firestore:"name"`
	Address       string    `firestore:"address"`
	Img           string    `firestore:"img"`
	Entry         string    `firestore:"entry"`
	SiteURL       string    `firestore:"siteUrl"`
	Lat           float64   `firestore:"lat"`
	Lng           float64   `firestore:"lng"`
	ParentID      string    `firestore:"parentId"`
	Source        string    `firestore:"source"`
	SubIdentifier string    `firestore:"subIdentifier"`
}

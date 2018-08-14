package repository

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/andoshin11/go-crawler-example/src/types"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/api/iterator"
)

const collection = "crawlerExample"

// MuseumRepository interface
type MuseumRepository interface {
	GetAll(ctx context.Context) ([]*types.Museum, error)
	GetBySubIdentifier(ctx context.Context, subIdentifier string) (*types.Museum, error)
	AddItem(ctx context.Context, subIdentifier, source string) (err error)
	UpdateItem(ctx context.Context, identifier string, museum *types.Museum) (err error)
}

type museumRepository struct {
	Client *firestore.Client
}

// NewMuseumRepository return struct
func NewMuseumRepository(Client *firestore.Client) MuseumRepository {
	return &museumRepository{Client}
}

func (r *museumRepository) GetAll(ctx context.Context) ([]*types.Museum, error) {
	iter := r.Client.Collection(collection).Documents(ctx)

	museums := make([]*types.Museum, 10)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}

		museum := types.Museum{}
		doc.DataTo(&museum)
		museums = append(museums, &museum)
	}

	return museums, nil
}

func (r *museumRepository) GetBySubIdentifier(ctx context.Context, subIdentifier string) (*types.Museum, error) {
	iter := r.Client.Collection(collection).Where("subIdentifier", "==", subIdentifier).Documents(ctx)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		museum := types.Museum{}
		doc.DataTo(&museum)

		return &museum, err
	}

	return nil, errors.New("Item not found")
}

func (r *museumRepository) AddItem(ctx context.Context, subIdentifier, source string) (err error) {
	identifier := uuid.NewV4().String()
	t := time.Now()
	_, err = r.Client.Collection(collection).Doc(identifier).Set(ctx, types.Museum{
		CreatedAt:     t,
		SubIdentifier: subIdentifier,
		Identifier:    identifier,
		Source:        source,
	})

	return
}

func (r *museumRepository) UpdateItem(ctx context.Context, identifier string, museum *types.Museum) (err error) {
	t := time.Now()

	fmt.Printf("Success %#v\n", museum.Name)
	fmt.Println(reflect.ValueOf(museum).Kind() != reflect.Map)
	fmt.Printf("Success %#v\n", identifier)

	_, err = r.Client.Collection(collection).Doc(identifier).Set(ctx, StructToMap(types.Museum{
		Name:      museum.Name,
		Address:   museum.Address,
		Img:       museum.Img,
		Entry:     museum.Entry,
		SiteURL:   museum.SiteURL,
		UpdatedAt: t,
		Lat:       museum.Lat,
		Lng:       museum.Lng,
	}), firestore.MergeAll)

	return
}

func StructToMap(data interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	elem := reflect.ValueOf(data).Elem()
	size := elem.NumField()

	for i := 0; i < size; i++ {
		field := elem.Type().Field(i).Name
		value := elem.Field(i).Interface()
		result[field] = value
	}

	return result
}

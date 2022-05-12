package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"sync"

	models "gql/domain/model"
	"gql/graph/model"
)

type Resolver struct {
	Repo        models.RepositoryInterface
	subscribers map[string]chan<- *model.Photo
	mutex       sync.Mutex
	PhotoID     int64
}

func NewResolver(repo models.RepositoryInterface, photoID int64) *Resolver {
	return &Resolver{
		Repo:        repo,
		subscribers: map[string]chan<- *model.Photo{},
		mutex:       sync.Mutex{},
		PhotoID:     photoID,
	}
}

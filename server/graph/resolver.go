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
	Repo             models.RepositoryInterface
	photoSubscribers map[string]chan<- *model.Photo
	userSubscribers  map[string]chan<- []*model.User
	mutex            sync.Mutex
	PhotoID          int64
}

func NewResolver(repo models.RepositoryInterface, photoID int64) *Resolver {
	return &Resolver{
		Repo:             repo,
		photoSubscribers: map[string]chan<- *model.Photo{},
		userSubscribers:  map[string]chan<- []*model.User{},
		mutex:            sync.Mutex{},
		PhotoID:          photoID,
	}
}

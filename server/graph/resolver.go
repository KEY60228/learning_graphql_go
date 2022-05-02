package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"gql/domain/model"
)

type Resolver struct {
	Repo    model.RepositoryInterface
	PhotoID int64
}

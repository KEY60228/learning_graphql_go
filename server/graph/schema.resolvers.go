package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"gql/graph/generated"
	"gql/graph/model"
	"strconv"
)

func (r *mutationResolver) PostPhoto(ctx context.Context, input model.PostPhotoInput) (*model.Photo, error) {
	ID++
	photo := &model.Photo{
		ID:          strconv.Itoa(ID),
		Name:        input.Name,
		URL:         fmt.Sprintf("https://example.com/img/%d.jpg", ID),
		Description: input.Description,
		Category:    *input.Category,
	}
	Photos = append(Photos, photo)
	return photo, nil
}

func (r *queryResolver) TotalPhotos(ctx context.Context) (int, error) {
	return len(Photos), nil
}

func (r *queryResolver) AllPhotos(ctx context.Context) ([]*model.Photo, error) {
	return Photos, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
var ID int
var Photos []*model.Photo

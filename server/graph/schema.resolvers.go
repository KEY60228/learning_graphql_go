package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"gql/domain/service"
	"gql/graph/generated"
	"gql/graph/model"
)

func (r *mutationResolver) PostPhoto(ctx context.Context, input model.PostPhotoInput) (*model.Photo, error) {
	var description string
	if input.Description != nil {
		description = *input.Description
	} else {
		description = ""
	}

	taggedUserIDs := make([]string, len(input.TaggedUserIDs))
	for i, id := range input.TaggedUserIDs {
		taggedUserIDs[i] = id
	}

	r.PhotoID++

	photo, err := service.NewService(r.Repo).PostPhoto(
		ctx,
		r.PhotoID,
		input.Name,
		description,
		input.Category.String(),
		input.PostedByUserID,
		taggedUserIDs,
	)
	if err != nil {
		return nil, err
	}

	return photo, nil
}

func (r *queryResolver) TotalPhotos(ctx context.Context) (int, error) {
	count, err := service.NewService(r.Repo).TotalPhotos(ctx)
	if err != nil {
		return -1, err
	}
	return count, nil
}

func (r *queryResolver) AllPhotos(ctx context.Context) ([]*model.Photo, error) {
	photos, err := service.NewService(r.Repo).AllPhotos(ctx)
	if err != nil {
		return nil, err
	}
	return photos, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

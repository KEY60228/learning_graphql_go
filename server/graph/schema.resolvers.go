package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"gql/domain/service"
	"gql/graph/generated"
	"gql/graph/model"
	"gql/middleware"
	"gql/support"
)

func (r *mutationResolver) PostPhoto(ctx context.Context, input model.PostPhotoInput) (*model.Photo, error) {
	user, _ := ctx.Value(middleware.UserCtxKey).(*model.User)
	if user == nil {
		return nil, errors.New("unauthorized")
	}

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
		user.GithubLogin,
		taggedUserIDs,
	)
	if err != nil {
		return nil, err
	}

	return photo, nil
}

func (r *mutationResolver) GithubAuth(ctx context.Context, code string) (*model.AuthPayload, error) {
	data, accessToken := support.AuthorizeWithGitHub(support.GITHUB_CLIENT_ID, support.GITHUB_CLIENT_SECRETS, code)

	if data.Message != "" {
		return nil, errors.New(data.Message)
	}

	user, err := model.NewUser(data.Login, data.Name, data.AvatarUrl)
	if err != nil {
		return nil, err
	}

	service.NewService(r.Repo).UpdateUser(user, accessToken)

	auth, err := model.NewAuthPayload(accessToken, user)
	if err != nil {
		return nil, err
	}

	return auth, nil
}

func (r *mutationResolver) AddFakeUsers(ctx context.Context, count int) ([]*model.User, error) {
	users, tokens, err := support.AddFakeUsers(count)
	if err != nil {
		return nil, err
	}

	serv := service.NewService(r.Repo)
	for i, user := range users {
		serv.PostUser(user.GithubLogin, user.Name, user.Avatar, tokens[i])
	}

	return users, nil
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

func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	user, ok := ctx.Value(middleware.UserCtxKey).(*model.User)
	if !ok {
		return nil, errors.New("bad request")
	}
	return user, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

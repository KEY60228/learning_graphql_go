package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
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
	fmt.Println(*user)

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

	r.mutex.Lock()
	for _, ch := range r.photoSubscribers {
		ch <- photo
	}
	r.mutex.Unlock()

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

	r.mutex.Lock()
	for _, ch := range r.userSubscribers {
		ch <- users
	}
	r.mutex.Unlock()

	return users, nil
}

func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	user, ok := ctx.Value(middleware.UserCtxKey).(*model.User)
	if !ok {
		return nil, errors.New("bad request")
	}
	return user, nil
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

func (r *queryResolver) TotalUsers(ctx context.Context) (int, error) {
	count, err := service.NewService(r.Repo).TotalUsers(ctx)
	if err != nil {
		return -1, err
	}
	return count, err
}

func (r *queryResolver) AllUsers(ctx context.Context) ([]*model.User, error) {
	users, err := service.NewService(r.Repo).AllUsers(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *subscriptionResolver) NewUsers(ctx context.Context, githubLogin string) (<-chan []*model.User, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, ok := r.userSubscribers[githubLogin]; ok {
		err := fmt.Errorf("`%s` has already been subscribed", githubLogin)
		return nil, err
	}

	ch := make(chan []*model.User, 1)
	r.userSubscribers[githubLogin] = ch

	go func() {
		<-ctx.Done()
		r.mutex.Lock()
		delete(r.userSubscribers, githubLogin)
		r.mutex.Unlock()
	}()

	return ch, nil
}

func (r *subscriptionResolver) NewPhoto(ctx context.Context, githubLogin string) (<-chan *model.Photo, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, ok := r.photoSubscribers[githubLogin]; ok {
		err := fmt.Errorf("`%s` has already been subscribed", githubLogin)
		return nil, err
	}

	ch := make(chan *model.Photo, 1)
	r.photoSubscribers[githubLogin] = ch

	go func() {
		<-ctx.Done()
		r.mutex.Lock()
		delete(r.photoSubscribers, githubLogin)
		r.mutex.Unlock()
	}()

	return ch, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }

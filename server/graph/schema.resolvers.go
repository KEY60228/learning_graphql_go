package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"gql/graph/generated"
	"gql/graph/model"
)

func (r *mutationResolver) PostPhoto(ctx context.Context, input model.PostPhotoInput) (*model.Photo, error) {
	ID++

	photo := &model.Photo{
		ID:          strconv.Itoa(ID),
		URL:         fmt.Sprintf("https://example.com/img/%d.jpg", ID),
		Name:        input.Name,
		Description: input.Description,
		Category:    *input.Category,
		PostedBy:    getUserByGithubLogin("mHattrup"),
		Created:     model.DateTime(time.Now()),
	}
	Photos = append(Photos, photo)

	tag := []struct {
		PhotoID string
		UserID  string
	}{{
		PhotoID: strconv.Itoa(ID),
		UserID:  "mHattrup",
	}, {
		PhotoID: strconv.Itoa(ID),
		UserID:  "gPlake",
	}}
	Tags = append(Tags, tag...)

	return photo, nil
}

func (r *queryResolver) TotalPhotos(ctx context.Context) (int, error) {
	return len(Photos), nil
}

func (r *queryResolver) AllPhotos(ctx context.Context) ([]*model.Photo, error) {
	photos := make([]*model.Photo, len(Photos))
	for i, photo := range Photos {
		copy := *photo
		for _, tag := range Tags {
			if tag.PhotoID == copy.ID {
				copy.TaggedUsers = append(copy.TaggedUsers, getUserByGithubLogin(tag.UserID))
			}
		}
		photos[i] = &copy
	}
	return photos, nil
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
var Users []*model.User
var Photos []*model.Photo
var Tags []struct {
	PhotoID string
	UserID  string
}

func init() {
	Users = []*model.User{
		{GithubLogin: "mHattrup", Name: "Mike Hattrup"},
		{GithubLogin: "gPlake", Name: "Glen Plake"},
		{GithubLogin: "sSchmidt", Name: "Scot Schmidt"},
	}
}
func init() {
	Photos = []*model.Photo{
		{
			ID:          "1",
			Name:        "Dropping the Heart Chute",
			Description: toPtr("The heart chute is one of my favorite chutes"),
			Category:    model.PhotoCategoryAction,
			PostedBy:    getUserByGithubLogin("gPlake"),
			Created:     model.DateTime(stringToTime("1977/3/28")),
		},
		{
			ID:       "2",
			Name:     "Enjoying the sunshine",
			Category: model.PhotoCategorySelfie,
			PostedBy: getUserByGithubLogin("sSchmidt"),
			Created:  model.DateTime(stringToTime("1985/1/24")),
		},
		{
			ID:          "3",
			Name:        "Gunbarrel 25",
			Description: toPtr("25 laps on gunbarrel today"),
			Category:    model.PhotoCategoryLandscape,
			PostedBy:    getUserByGithubLogin("sSchmidt"),
			Created:     model.DateTime(stringToTime("2018/04/15")),
		},
	}
	ID += len(Photos)
}
func init() {
	Tags = []struct {
		PhotoID string
		UserID  string
	}{
		{PhotoID: "1", UserID: "gPlake"},
		{PhotoID: "2", UserID: "sSchmidt"},
		{PhotoID: "2", UserID: "mHattrup"},
		{PhotoID: "2", UserID: "gPlake"},
	}
}
func toPtr(s string) *string {
	return &s
}
func stringToTime(s string) time.Time {
	t, _ := time.Parse("2006/1/2", s)
	return t
}
func getUserByGithubLogin(s string) *model.User {
	for _, user := range Users {
		if user.GithubLogin == s {
			return user
		}
	}
	return nil
}

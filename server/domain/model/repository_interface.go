package model

import (
	"gql/graph/model"
	"time"
)

type RepositoryInterface interface {
	PostPhoto(int64, string, string, string, string, string, []string, time.Time) error
	TotalPhotos() int
	AllPhotos() []*model.Photo
	UserByID(string) *model.User
	UsersByIDs([]string) []*model.User
}

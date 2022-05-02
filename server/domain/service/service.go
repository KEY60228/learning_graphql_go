package service

import (
	"context"
	"fmt"
	"time"

	models "gql/domain/model"
	"gql/graph/model"
)

type Service struct {
	Repo models.RepositoryInterface
}

func NewService(repo models.RepositoryInterface) *Service {
	return &Service{Repo: repo}
}

func (s *Service) PostPhoto(
	ctx context.Context,
	photoID int64,
	name string,
	description string,
	category string,
	postedUserID string,
	taggedUserIDs []string,
) (*model.Photo, error) {
	user := s.UserByID(postedUserID)
	users := s.UsersByIDs(taggedUserIDs)
	now := time.Now()

	err := s.Repo.PostPhoto(photoID, fmt.Sprintf("https://example.com/img/%d.jpg", photoID), name, description, category, postedUserID, taggedUserIDs, now)
	if err != nil {
		return nil, err
	}

	photo, err := model.NewPhoto(photoID, fmt.Sprintf("https://example.com/img/%d.jpg", photoID), name, description, category, user, users, now)
	if err != nil {
		return nil, err
	}

	return photo, nil
}

func (s *Service) TotalPhotos(ctx context.Context) (int, error) {
	return s.Repo.TotalPhotos(), nil
}

func (s *Service) AllPhotos(ctx context.Context) ([]*model.Photo, error) {
	photos := s.Repo.AllPhotos()
	return photos, nil
}

func (s *Service) UserByID(id string) *model.User {
	return s.Repo.UserByID(id)
}

func (s *Service) UsersByIDs(ids []string) []*model.User {
	return s.Repo.UsersByIDs(ids)
}

func (s *Service) UpdateUser(user *model.User, accessToken string) error {
	s.Repo.UpdateUser(user.GithubLogin, user.Avatar, accessToken)
	return nil
}

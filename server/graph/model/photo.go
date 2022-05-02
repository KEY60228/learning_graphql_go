package model

import (
	"errors"
	"strconv"
	"time"
)

func NewPhoto(id int64, url string, name string, description string, category string, postedBy *User, taggedUsers []*User, created time.Time) (*Photo, error) {
	if !PhotoCategory(category).IsValid() {
		return nil, errors.New("category error")
	}

	return &Photo{
		ID:          strconv.Itoa(int(id)),
		URL:         url,
		Name:        name,
		Description: &description,
		Category:    PhotoCategory(category),
		PostedBy:    postedBy,
		TaggedUsers: taggedUsers,
		Created:     DateTime(created),
	}, nil
}

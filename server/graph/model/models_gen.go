// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type AuthPayload struct {
	Token string `json:"token"`
	User  *User  `json:"user"`
}

type Photo struct {
	ID          string        `json:"id"`
	URL         string        `json:"url"`
	Name        string        `json:"name"`
	Description *string       `json:"description"`
	Category    PhotoCategory `json:"category"`
	PostedBy    *User         `json:"postedBy"`
	TaggedUsers []*User       `json:"taggedUsers"`
	Created     DateTime      `json:"created"`
}

type PostPhotoInput struct {
	Name          string         `json:"name"`
	Description   *string        `json:"description"`
	Category      *PhotoCategory `json:"category"`
	TaggedUserIDs []string       `json:"taggedUserIDs"`
}

type User struct {
	GithubLogin string `json:"githubLogin"`
	Name        string `json:"name"`
	Avatar      string `json:"avatar"`
}

type PhotoCategory string

const (
	PhotoCategorySelfie    PhotoCategory = "SELFIE"
	PhotoCategoryPortrait  PhotoCategory = "PORTRAIT"
	PhotoCategoryAction    PhotoCategory = "ACTION"
	PhotoCategoryLandscape PhotoCategory = "LANDSCAPE"
	PhotoCategoryGraphic   PhotoCategory = "GRAPHIC"
)

var AllPhotoCategory = []PhotoCategory{
	PhotoCategorySelfie,
	PhotoCategoryPortrait,
	PhotoCategoryAction,
	PhotoCategoryLandscape,
	PhotoCategoryGraphic,
}

func (e PhotoCategory) IsValid() bool {
	switch e {
	case PhotoCategorySelfie, PhotoCategoryPortrait, PhotoCategoryAction, PhotoCategoryLandscape, PhotoCategoryGraphic:
		return true
	}
	return false
}

func (e PhotoCategory) String() string {
	return string(e)
}

func (e *PhotoCategory) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PhotoCategory(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PhotoCategory", str)
	}
	return nil
}

func (e PhotoCategory) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

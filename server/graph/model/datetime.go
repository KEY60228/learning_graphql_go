package model

import (
	"fmt"
	"io"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

// DateTime represents "DateTime" GraphQL custom scalar
type DateTime time.Time

// MarshalGQL implements graphql.Marshaler interface
func (dt DateTime) MarshalGQL(w io.Writer) {
	graphql.MarshalTime(time.Time(dt)).MarshalGQL(w)
}

// UnmarshalGQL implements graphql.Marshaler interface
func (dt *DateTime) UnmarshalGQL(v interface{}) error {
	t, err := graphql.UnmarshalTime(v)
	if err != nil {
		return fmt.Errorf("invalid DateTime: %v, %v", v, err)
	}

	*dt = DateTime(t)
	return nil
}

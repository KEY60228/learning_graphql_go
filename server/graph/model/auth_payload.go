package model

func NewAuthPayload(token string, user *User) (*AuthPayload, error) {
	return &AuthPayload{
		Token: token,
		User:  user,
	}, nil
}

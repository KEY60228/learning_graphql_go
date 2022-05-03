package model

func NewUser(githubLogin string, name string, avatar string) (*User, error) {
	return &User{
		GithubLogin: githubLogin,
		Name:        name,
		Avatar:      avatar,
	}, nil
}

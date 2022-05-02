package model

func NewUser(githubLogin string, name string) (*User, error) {
	return &User{
		GithubLogin: githubLogin,
		Name:        name,
	}, nil
}

package support

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"gql/graph/model"
)

type GitHubAccessTokenResponse struct {
	Accesstoken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

type GitHubData struct {
	Message   string `json:"message"`
	AvatarUrl string `json:"avatar_url"`
	Login     string `json:"login"`
	Name      string `json:"name"`
}

type FakeResult struct {
	FakeUsers []*FakeUser `json:"results"`
}

type FakeUser struct {
	Name    FakeUserName    `json:"name"`
	Login   FakeUserLogin   `json:"login"`
	Picture FakeUserPicture `json:"picture"`
}

type FakeUserName struct {
	FirstName string `json:"first"`
	LastName  string `json:"last"`
}

type FakeUserLogin struct {
	GithubLogin string `json:"username"`
	AccessToken string `json:"sha1"`
}

type FakeUserPicture struct {
	Avatar string `json:"thumbnail"`
}

func AddFakeUsers(count int) ([]*model.User, []string, error) {
	res, err := http.Get(fmt.Sprintf("https://randomuser.me/api/?results=%d", count))
	if err != nil {
		return nil, nil, err
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, nil, err
	}

	var results FakeResult
	json.Unmarshal(resBody, &results)

	users := make([]*model.User, len(results.FakeUsers))
	accessTokens := make([]string, len(results.FakeUsers))
	for i, fakeUser := range results.FakeUsers {
		users[i], err = model.NewUser(fakeUser.Login.GithubLogin, fakeUser.Name.LastName+fakeUser.Name.LastName, fakeUser.Picture.Avatar)
		if err != nil {
			return nil, nil, err
		}
		accessTokens[i] = fakeUser.Login.AccessToken
	}

	return users, accessTokens, nil
}

func AuthorizeWithGitHub(clientID string, clientSecret string, code string) (GitHubData, string) {
	accessToken := requestGitHubToken(clientID, clientSecret, code)
	githubUser := requestGitHubUserAccount(accessToken)
	return githubUser, accessToken
}

func requestGitHubToken(clientID string, clientSecret string, code string) string {
	requestBodyMap := map[string]string{
		"client_id":     clientID,
		"client_secret": clientSecret,
		"code":          code,
	}
	requestJson, err := json.Marshal(requestBodyMap)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("POST", "https://github.com/login/oauth/access_token", bytes.NewBuffer(requestJson))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var githubRes GitHubAccessTokenResponse
	json.Unmarshal(resBody, &githubRes)

	return githubRes.Accesstoken
}

func requestGitHubUserAccount(accessToken string) GitHubData {
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("token %s", accessToken))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var githubData GitHubData
	json.Unmarshal(resBody, &githubData)

	return githubData
}

package support

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

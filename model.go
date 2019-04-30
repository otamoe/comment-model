package model

import (
	"net/http"

	authModel "github.com/otamoe/auth-model"
)

var (
	APIOrigin  string
	PostOrigin string
	PushURL    string
	Client     *http.Client
)

func Config(apiOrigin, postOrigin, pushURL string) {
	APIOrigin = apiOrigin
	PostOrigin = postOrigin
	PushURL = pushURL
}

func Start() {
	Client = authModel.GetClientCredentials([]string{"application:all", "post:all"})
}

//
// func UpdatePost() {
// 	// Client.Do()
// }

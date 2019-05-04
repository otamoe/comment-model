package model

import (
	"testing"

	authModel "github.com/otamoe/auth-model"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetLevel(logrus.TraceLevel)
	authModel.Config("http://auth.auth.de:8080", "http://user.auth.de:8080", "5cbc45ef11ca2b6e6c6f139e", "wwwwwww")
	Config("http://api.comment.de:8110", "http://application.comment.de:8110", "http://qq.com")
	Start()
}

func TestPost(t *testing.T) {
	post := &Post{}
	post.URI = "/test/xxxx"
	if err := post.Save(); err != nil {
		t.Error(err)
	}
	t.Log(post)
}

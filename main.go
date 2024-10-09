package main

import (
	"context"
	"fmt"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

func main() {
	/**
		GITHUB_CLIENT_ID=6903aa8547a474fb9202
	GITHUB_CLIENT_SECRET=3fda67d3f8ca9aa0ba663d05f294f7c1915278b5
	GITHUB_AUTH_REDIRECT=http://api.opensource.fourm/v1/auth/callback
	**/
	conf := &oauth2.Config{
		ClientID:     "6903aa8547a474fb9202",
		ClientSecret: "3fda67d3f8ca9aa0ba663d05f294f7c1915278b5",
		Scopes:       []string{},
		Endpoint:     github.Endpoint,
		RedirectURL:  "http://api.opensource.fourm/v1/auth/callback",
	}
	code := "9d4f2dc1d916cdae9d53"

	token, err := conf.Exchange(context.Background(), code)
	fmt.Println(token, err)
}

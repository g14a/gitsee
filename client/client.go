package client

import (
	"context"
	"github.com/shurcooL/githubv4"
	_ "github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
	"sync"
)

var (
	GHClient  githubv4.Client
	GHContext context.Context
	once      sync.Once
	ghToken   = "12d2647357241f80e2959a0c448efa299b803511"
)

func init() {
	once.Do(func() {
		GHClient, GHContext = getGHClient(ghToken)
	})
}

func getGHClient(token string) (githubv4.Client, context.Context) {
	ghContext := context.Background()
	tokenService := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tokenClient := oauth2.NewClient(ghContext, tokenService)

	return *githubv4.NewClient(tokenClient), ghContext
}

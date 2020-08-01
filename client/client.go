package client

import (
	"context"
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
	"os"
	"sync"
)

var (
	GHClient  githubv4.Client
	GHContext context.Context
	once      sync.Once
)

func init() {
	once.Do(func() {
		ghToken := os.Getenv("GHTOKEN")
		fmt.Println(ghToken, "is ghtoken")
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

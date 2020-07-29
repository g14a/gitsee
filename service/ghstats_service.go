package service

import (
	"fmt"
	"github.com/shurcooL/githubv4"
	"gitsee/client"
)

func GetRepoStats(owner string) {
	err := client.GHClient.Query(client.GHContext, &query, nil)
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(query.Viewer.Login)
	fmt.Println(query.Viewer.CreatedAt)
}

var query struct {
	Viewer struct {
		Login githubv4.String
		CreatedAt githubv4.DateTime
	}
}
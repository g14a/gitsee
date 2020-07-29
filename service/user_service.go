package service

import (
	"fmt"
	"github.com/shurcooL/githubv4"
	"gitsee/client"
)

func UserDetails(user string) {
	
	var query struct {
		User struct {
			Name githubv4.String
			CreatedAt githubv4.DateTime
			AvatarURL githubv4.String
			Location githubv4.String
			Bio githubv4.String
			Followers struct {
				TotalCount githubv4.Int
			}
		} `graphql:"user(login: $user)"`
	}
	
	variables := map[string]interface{} {
		"user": githubv4.String(user),
	}
	
	err := client.GHClient.Query(client.GHContext, &query, variables)
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(query)
}

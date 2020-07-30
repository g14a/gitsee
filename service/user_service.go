package service

import (
	"github.com/shurcooL/githubv4"
	"gitsee/client"
	"log"
)

func UserDetails(user string) (map[string]interface{}, error) {

	variables := map[string]interface{}{
		"user": githubv4.String(user),
	}

	err := client.GHClient.Query(client.GHContext, &UserQuery, variables)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	mapUser := UserQuery.User

	return map[string]interface{}{
		"name":       mapUser.Name,
		"created_at": mapUser.CreatedAt,
		"bio":        mapUser.Bio,
		"location":   mapUser.Location,
		"avatar_url": mapUser.AvatarURL,
		"followers":  mapUser.Followers.TotalCount,
	}, nil

}

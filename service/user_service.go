package service

import (
	"encoding/json"
	"fmt"
	"github.com/shurcooL/githubv4"
	"gitsee/client"
)

func UserDetails(user string) {

	variables := map[string]interface{}{
		"user": githubv4.String(user),
	}

	err := client.GHClient.Query(client.GHContext, &UserQuery, variables)
	if err != nil {
		fmt.Println(err)
	}

	mapUser := UserQuery.User

	jsonResponse := map[string]interface{}{
		"name":       mapUser.Name,
		"created_at": mapUser.CreatedAt,
		"bio":        mapUser.Bio,
		"location":   mapUser.Location,
		"avatar_url": mapUser.AvatarURL,
		"followers":  mapUser.Followers.TotalCount,
	}

	bytes, _ := json.Marshal(jsonResponse)
	fmt.Println(string(bytes))
}

package service

import (
	"gitsee/cache"
	"gitsee/client"
	"log"

	"github.com/dustin/go-humanize"
	"github.com/shurcooL/githubv4"
)

func UserDetails(user string) (interface{}, error) {

	if userDetails, ok := cache.Get(user); ok {
		log.Println("Got user details from cache")
		return userDetails, nil
	}

	variables := map[string]interface{}{
		"user": githubv4.String(user),
	}

	err := client.GHClient.Query(client.GHContext, &UserQuery, variables)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	mapUser := UserQuery.User
	jsonResponse := map[string]interface{}{
		"username":   mapUser.Login,
		"name":       mapUser.Name,
		"created_at": "Joined " + humanize.Time(mapUser.CreatedAt.Time),
		"location":   mapUser.Location,
		"avatar_url": mapUser.AvatarURL,
		"followers":  humanize.Comma(int64(mapUser.Followers.TotalCount)),
		"url":        mapUser.URL,
	}

	if cache.Set(user, jsonResponse) {
		log.Println("Set ", user, "details in Cache")
	}

	return jsonResponse, nil
}

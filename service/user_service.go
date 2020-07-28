package service

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"gitsee/client"
)

func UserDetails(user string) (map[string]interface{}, error) {
	ghUser, _, err := client.GHClient.Users.Get(client.GHContext, user)
	
	if err != nil {
		fmt.Printf("Problem in getting user information %v\n", err)
		return nil, err
	}
	
	return map[string]interface{}{
		"user": map[string]interface{}{
			"name":      ghUser.GetName(),
			"joined":    "Joined GitHub " + humanize.Time(ghUser.GetCreatedAt().Time),
			"location":  ghUser.GetLocation(),
			"avatar":    ghUser.GetAvatarURL(),
			"url":       ghUser.GetHTMLURL(),
			"followers": ghUser.GetFollowers(),
		},
	}, nil
}

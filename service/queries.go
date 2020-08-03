package service

import (
	"github.com/shurcooL/githubv4"
)

// UserQuery start
var UserQuery struct {
	User struct {
		Login     githubv4.String
		Name      githubv4.String
		CreatedAt githubv4.DateTime
		AvatarURL githubv4.URI
		Location  githubv4.String
		URL       githubv4.URI
		Followers struct {
			TotalCount githubv4.Int
		}
	} `graphql:"user(login: $user)"`
}

// UserQuery end

// ForksStarsLanguagesQuery start
var ForksStarsLanguagesQuery struct {
	User struct {
		Repositories struct {
			Nodes []Nodes
		} `graphql:"repositories(first: $repoCount, ownerAffiliations: OWNER)"`
	} `graphql:"user(login: $user)"`
}

type Nodes struct {
	PrimaryLanguage struct {
		Name string
	}
	Watchers struct {
		TotalCount int
	}
	StarGazers struct {
		TotalCount int
	} `graphql:"stargazers"`
	Name      string
	ForkCount int
	Languages struct {
		TotalCount int
		Nodes      []Language
	} `graphql:"languages(first: $languageCount)"`
}

type Commit struct {
	History struct {
		TotalCount githubv4.Int
	}
}

type Language struct {
	Name string
}

// ForksStarsLanguagesQuery end

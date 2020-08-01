package service

import (
	"github.com/shurcooL/githubv4"
)

// UserQuery start
var UserQuery struct {
	User struct {
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
		Name githubv4.String
	}
	Watchers struct {
		TotalCount githubv4.Int
	}
	StarGazers struct {
		TotalCount githubv4.Int
	} `graphql:"stargazers"`
	Name      githubv4.String
	ForkCount githubv4.Int
	Languages struct {
		TotalCount githubv4.Int
		Nodes      []Language
	} `graphql:"languages(first: $languageCount)"`
}

type Commit struct {
	History struct {
		TotalCount githubv4.Int
	}
}

type Language struct {
	Name githubv4.String
}

// ForksStarsLanguagesQuery end

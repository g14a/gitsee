package models

type AbstractRepo struct {
	RepoName        string
	Languages       map[string]int
	UserCommitCount int // Total commit count of the user who is the owner as well
	StarCount       int
	ForksCount      int
}

type AbsoluteResponse struct {
	UserDetails          map[string]interface{}
	FrequencyOfLanguages map[string]int
	ReposStars           map[string]interface{}
	ReposForks           map[string]interface{}
	ReposCommits         map[string]interface{}
}

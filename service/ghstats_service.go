package service

import (
	"fmt"
	"github.com/shurcooL/githubv4"
	"gitsee/client"
)

var (
	ReposForks          map[string]interface{}
	ReposStars          map[string]interface{}
	LanguageFrequencies map[string]int
	PrimaryLanguages    map[string]int
	PrimaryLanguageStars map[string]int
)

func ForksStarsLanguages(user string, repoCount, languageCount int) {
	variables := map[string]interface{}{
		"user":          githubv4.String(user),
		"repoCount":     githubv4.Int(repoCount),
		"languageCount": githubv4.Int(languageCount),
	}

	err := client.GHClient.Query(client.GHContext, &ForksStarsLanguagesQuery, variables)
	if err != nil {
		fmt.Println(err)
	}

	reposForks := make(map[string]interface{})

	for _, v := range ForksStarsLanguagesQuery.User.Repositories.Nodes {
		if v.ForkCount > 0 {
			reposForks[string(v.Name)] = v.ForkCount
		}
	}

	reposStars := make(map[string]interface{})

	for _, v := range ForksStarsLanguagesQuery.User.Repositories.Nodes {
		if v.StarGazers.TotalCount > 0 {
			reposStars[string(v.Name)] = v.StarGazers.TotalCount
		}
	}

	ReposForks = reposForks
	ReposStars = reposStars

	languageFrequencies := make(map[string]int)

	for _, v := range ForksStarsLanguagesQuery.User.Repositories.Nodes {
		for _, repo := range v.Languages.Nodes {
			languageFrequencies[string(repo.Name)] += 1
		}
	}

	primaryLanguages := make(map[string]int)

	for _, v := range ForksStarsLanguagesQuery.User.Repositories.Nodes {
		primaryLanguages[string(v.PrimaryLanguage.Name)] += 1
	}
	
	primaryLanguageStars := make(map[string]int)
	
	for _, v := range ForksStarsLanguagesQuery.User.Repositories.Nodes {
		if v.StarGazers.TotalCount > 0 && len(v.PrimaryLanguage.Name) > 0 {
			primaryLanguageStars[string(v.PrimaryLanguage.Name)] += int(v.StarGazers.TotalCount)
		}
	}
	
	LanguageFrequencies = languageFrequencies
}

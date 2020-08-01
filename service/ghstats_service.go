package service

import (
	"github.com/shurcooL/githubv4"
	"gitsee/cache"
	"gitsee/client"
	"log"
	"time"
)

var (
	ReposForks           = make(map[string]interface{})
	ReposStars           = make(map[string]interface{})
	LanguageFrequencies  = make(map[string]int)
	PrimaryLanguages     = make(map[string]int)
	PrimaryLanguageStars = make(map[string]int)
)

func ForksStarsLanguages(user string, repoCount, languageCount int) error {
	variables := map[string]interface{}{
		"user":          githubv4.String(user),
		"repoCount":     githubv4.Int(repoCount),
		"languageCount": githubv4.Int(languageCount),
	}

	err := client.GHClient.Query(client.GHContext, &ForksStarsLanguagesQuery, variables)
	if err != nil {
		log.Println(err)
		return err
	}

	repoForks := make(map[string]interface{})

	for _, v := range ForksStarsLanguagesQuery.User.Repositories.Nodes {
		if v.ForkCount > 0 {
			repoForks[string(v.Name)] = v.ForkCount
		}
	}

	ReposForks = repoForks

	if cache.Set(user+"RepoForks", repoForks) {
		log.Println(user + "RepoForks added to cache")
	}
	
	repoStars := make(map[string]interface{})

	for _, v := range ForksStarsLanguagesQuery.User.Repositories.Nodes {
		if v.StarGazers.TotalCount > 0 {
			repoStars[string(v.Name)] = v.StarGazers.TotalCount
		}
	}

	ReposStars = repoStars

	if cache.Set(user+"RepoStars", repoStars) {
		log.Println(user + "RepoStars added to cache")
	}
	
	languageFrequencies := make(map[string]int)

	for _, v := range ForksStarsLanguagesQuery.User.Repositories.Nodes {
		for _, repo := range v.Languages.Nodes {
			languageFrequencies[string(repo.Name)] += 1
		}
	}

	LanguageFrequencies = languageFrequencies

	if cache.Set(user+"LanguageFrequencies", languageFrequencies) {
		log.Println(user + "LanguageFrequencies added to cache")
	}
	
	primaryLanguages := make(map[string]int)

	for _, v := range ForksStarsLanguagesQuery.User.Repositories.Nodes {
		if len(v.PrimaryLanguage.Name) > 0 {
			primaryLanguages[string(v.PrimaryLanguage.Name)] += 1
		}
	}

	if cache.Set(user+"PrimaryLanguages", primaryLanguages) {
		log.Println(user + "PrimaryLanguages added to cache")
	}
	
	PrimaryLanguages = primaryLanguages

	primaryLanguageStars := make(map[string]int)

	for _, v := range ForksStarsLanguagesQuery.User.Repositories.Nodes {
		if v.StarGazers.TotalCount > 0 && len(v.PrimaryLanguage.Name) > 0 {
			primaryLanguageStars[string(v.PrimaryLanguage.Name)] += int(v.StarGazers.TotalCount)
		}
	}

	if cache.Set(user+"PrimaryLanguageStars", primaryLanguageStars) {
		log.Println(user + "PrimaryLanguageStars added to cache")
	}
	
	PrimaryLanguageStars = primaryLanguageStars

	time.Sleep(time.Millisecond*10)
	
	return nil
}

func GetWantedStatFromCache(username, wantedStat string) (interface{}, bool) {
	// Get wanted stat from cache
	if result, ok := cache.Get(username + wantedStat); ok {
		return result, true
	} else {
		return nil, false
	}
}

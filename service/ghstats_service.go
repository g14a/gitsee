package service

import (
	"fmt"
	"github.com/shurcooL/githubv4"
	"gitsee/cache"
	"gitsee/client"
	"gitsee/color"
	"gitsee/utils"
	"log"
	"time"
)

var (
	ReposForks           = make(map[string]int)
	ReposStars           = make(map[string]int)
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

	repoForks := make(map[string]int)

	for _, v := range ForksStarsLanguagesQuery.User.Repositories.Nodes {
		if v.ForkCount > 0 {
			repoForks[v.Name] = v.ForkCount
		}
	}
	
	if len(repoForks) > 10 {
		sortedRepoForks := utils.GetSortedMap(repoForks)
		ReposForks = sortedRepoForks
		if cache.Set(user+"RepoForks", ReposForks) {
			log.Println(user + "RepoForks added to cache")
		}
	} else {
		ReposForks = repoForks
		if cache.Set(user+"RepoForks", ReposForks) {
			log.Println(user + "RepoForks added to cache")
		}
	}

	repoStars := make(map[string]int)

	for _, v := range ForksStarsLanguagesQuery.User.Repositories.Nodes {
		if v.StarGazers.TotalCount > 0 {
			repoStars[v.Name] = v.StarGazers.TotalCount
		}
	}

	if len(repoStars) >= 10 {
		sortedRepoStars := utils.GetSortedMap(repoStars)
		ReposStars = sortedRepoStars
		if cache.Set(user+"RepoStars", ReposStars) {
			log.Println(user + "RepoStars added to cache")
		}
	} else {
		ReposStars = repoStars
		if cache.Set(user+"RepoStars", ReposStars) {
			log.Println(user + "RepoStars added to cache")
		}
	}

	languageFrequencies := make(map[string]int)

	for _, v := range ForksStarsLanguagesQuery.User.Repositories.Nodes {
		for _, repo := range v.Languages.Nodes {
			languageFrequencies[repo.Name] += 1
		}
	}

	if len(languageFrequencies) >= 10 {
		sortedLanguagesFrequencies := utils.GetSortedMap(languageFrequencies)
		LanguageFrequencies = sortedLanguagesFrequencies
		if cache.Set(user+"LanguageFrequencies", sortedLanguagesFrequencies) {
			log.Println(user + "LanguageFrequencies added to cache")
		}
	} else {
		LanguageFrequencies = languageFrequencies
		if cache.Set(user+"LanguageFrequencies", languageFrequencies) {
			log.Println(user + "LanguageFrequencies added to cache")
		}
	}
	
	if len(languageFrequencies) > 0 {
		color.GetColorCodesForLanguages(user, LanguageFrequencies)
	}
	
	primaryLanguages := make(map[string]int)

	for _, v := range ForksStarsLanguagesQuery.User.Repositories.Nodes {
		if len(v.PrimaryLanguage.Name) > 0 {
			primaryLanguages[v.PrimaryLanguage.Name] += 1
		}
	}

	if len(primaryLanguages) >= 10 {
		fmt.Println(len(primaryLanguages), " is length of primary languages")
		sortedLanguages := utils.GetSortedMap(primaryLanguages)
		PrimaryLanguages = sortedLanguages
		if cache.Set(user+"PrimaryLanguages", sortedLanguages) {
			log.Println(user + "PrimaryLanguages added to cache")
		}
	} else {
		PrimaryLanguages = primaryLanguages
		if cache.Set(user+"PrimaryLanguages", primaryLanguages) {
			log.Println(user + "PrimaryLanguages added to cache")
		}
	}

	primaryLanguageStars := make(map[string]int)

	for _, v := range ForksStarsLanguagesQuery.User.Repositories.Nodes {
		if v.StarGazers.TotalCount > 0 && len(v.PrimaryLanguage.Name) > 0 {
			primaryLanguageStars[v.PrimaryLanguage.Name] += v.StarGazers.TotalCount
		}
	}

	if len(primaryLanguageStars) >= 10 {
		sortedLanguageStars := utils.GetSortedMap(primaryLanguageStars)
		PrimaryLanguageStars = sortedLanguageStars
		if cache.Set(user+"PrimaryLanguageStars", sortedLanguageStars) {
			log.Println(user + "PrimaryLanguageStars added to cache")
		}
	} else {
		PrimaryLanguageStars = primaryLanguageStars
		if cache.Set(user+"PrimaryLanguageStars", primaryLanguageStars) {
			log.Println(user + "PrimaryLanguageStars added to cache")
		}
	}

	time.Sleep(time.Millisecond * 10)

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

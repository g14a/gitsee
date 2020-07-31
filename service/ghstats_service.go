package service

import (
	"errors"
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

	if cache.RistrettoCache.SetWithTTL(user+"RepoForks", repoForks, 1, time.Second*5) {
		log.Println(user + "RepoForks added to cache")
	}

	time.Sleep(10 * time.Millisecond)

	repoStars := make(map[string]interface{})

	for _, v := range ForksStarsLanguagesQuery.User.Repositories.Nodes {
		if v.StarGazers.TotalCount > 0 {
			repoStars[string(v.Name)] = v.StarGazers.TotalCount
		}
	}

	ReposStars = repoStars

	if cache.RistrettoCache.SetWithTTL(user+"RepoStars", repoStars, 1, time.Second*5) {
		log.Println(user + "RepoStars added to cache")
	}

	time.Sleep(10 * time.Millisecond)

	languageFrequencies := make(map[string]int)

	for _, v := range ForksStarsLanguagesQuery.User.Repositories.Nodes {
		for _, repo := range v.Languages.Nodes {
			languageFrequencies[string(repo.Name)] += 1
		}
	}

	LanguageFrequencies = languageFrequencies

	if cache.RistrettoCache.SetWithTTL(user+"LanguageFrequencies", languageFrequencies, 1, time.Second*5) {
		log.Println(user + "LanguageFrequencies added to cache")
	}

	time.Sleep(10 * time.Millisecond)

	primaryLanguages := make(map[string]int)

	for _, v := range ForksStarsLanguagesQuery.User.Repositories.Nodes {
		if len(v.PrimaryLanguage.Name) > 0 {
			primaryLanguages[string(v.PrimaryLanguage.Name)] += 1
		}
	}

	if cache.RistrettoCache.SetWithTTL(user+"PrimaryLanguages", primaryLanguages, 1, time.Second*5) {
		log.Println(user + "PrimaryLanguages added to cache")
	}

	time.Sleep(10 * time.Millisecond)

	PrimaryLanguages = primaryLanguages

	primaryLanguageStars := make(map[string]int)

	for _, v := range ForksStarsLanguagesQuery.User.Repositories.Nodes {
		if v.StarGazers.TotalCount > 0 && len(v.PrimaryLanguage.Name) > 0 {
			primaryLanguageStars[string(v.PrimaryLanguage.Name)] += int(v.StarGazers.TotalCount)
		}
	}

	if cache.RistrettoCache.SetWithTTL(user+"PrimaryLanguageStars", primaryLanguageStars, 1, time.Second*5) {
		log.Println(user + "PrimaryLanguageStars added to cache")
	}

	time.Sleep(10 * time.Millisecond)

	PrimaryLanguageStars = primaryLanguageStars

	return nil
}

func GetWantedStatFromCache(username, wantedStat string) (interface{}, error) {
	// Get wanted stat from cache
	if result, ok := cache.RistrettoCache.Get(username + wantedStat); ok {
		return result, nil
	} else {
		return nil, errors.New("could not find stat in cache")
	}
}

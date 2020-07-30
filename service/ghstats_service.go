package service

import (
	"github.com/shurcooL/githubv4"
	"gitsee/client"
	"log"
	"sync"
)

var (
	ReposForks           = make(map[string]interface{})
	ReposStars           = make(map[string]interface{})
	LanguageFrequencies  = make(map[string]int)
	PrimaryLanguages     = make(map[string]int)
	PrimaryLanguageStars = make(map[string]int)
	once sync.Once
)

func GetAllStats(user string, repoCount, languageCount int)  {
	once.Do(func() {
		err := ForksStarsLanguages(user, repoCount, languageCount)
		if err != nil {
			log.Println(err)
			return
		}
	})
}

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

	for _, v := range ForksStarsLanguagesQuery.User.Repositories.Nodes {
		if v.ForkCount > 0 {
			ReposForks[string(v.Name)] = v.ForkCount
		}
	}
	
	for _, v := range ForksStarsLanguagesQuery.User.Repositories.Nodes {
		if v.StarGazers.TotalCount > 0 {
			ReposStars[string(v.Name)] = v.StarGazers.TotalCount
		}
	}
	
	
	for _, v := range ForksStarsLanguagesQuery.User.Repositories.Nodes {
		for _, repo := range v.Languages.Nodes {
			LanguageFrequencies[string(repo.Name)] += 1
		}
	}
	
	for _, v := range ForksStarsLanguagesQuery.User.Repositories.Nodes {
		if len(v.PrimaryLanguage.Name) > 0 {
			PrimaryLanguages[string(v.PrimaryLanguage.Name)] += 1
		}
	}
	
	for _, v := range ForksStarsLanguagesQuery.User.Repositories.Nodes {
		if v.StarGazers.TotalCount > 0 && len(v.PrimaryLanguage.Name) > 0 {
			PrimaryLanguageStars[string(v.PrimaryLanguage.Name)] += int(v.StarGazers.TotalCount)
		}
	}

	return nil
}

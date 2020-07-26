package service

import (
	"fmt"
	"github.com/google/go-github/github"
	"strings"
	"sync"
	"webs/client"
	"webs/models"
)

var (
	ReposOwnedByUser        []string
	DistinctLanguagesofUser []string
)

func GetUserStatsOfRepos(owner string) {
	stats, _, _ := client.GHClient.Repositories.List(client.GHContext, owner, &github.RepositoryListOptions{})

	// if err != nil {
	// 	fmt.Printf("Problem in getting repository information %v\n", err)
	// 	return
	// }

	repoStats := make([]models.AbstractRepo, 0)

	var wg sync.WaitGroup
	wg.Add(len(stats))

	GetRepoNames := func() {
		for _, v := range stats {
			v := v
			go func() {
				defer wg.Done()
				languageStats, _, _ := client.GHClient.Repositories.ListLanguages(client.GHContext, "g14a", v.GetName())
				
				repoStat := &models.AbstractRepo{
					RepoName:  v.GetName(),
					Languages: languageStats,
					StarCount: v.GetStargazersCount(),
				}
				
				repoStat.UserCommitCount = StatsOfContributor(v.GetName(), owner)
				repoStats = append(repoStats, *repoStat)
			}()
		}

		wg.Wait()
	}

	GetRepoNames()

	for _, v := range repoStats {
		ReposOwnedByUser = append(ReposOwnedByUser, v.RepoName)
	}

	fmt.Println(repoStats)
	
	GetDistinctLanguagesOfUsersRepos(repoStats)
}

func GetDistinctLanguagesOfUsersRepos(repoStats []models.AbstractRepo) {
	languages := models.NewLanguageSet()
	for _, v := range repoStats {
		for language, _ := range v.Languages {
			languages.Add(language)
		}
	}

	languageArray := make([]string, 0)
	for k, _ := range languages.Languages {
		languageArray = append(languageArray, k)
	}
	
	DistinctLanguagesofUser = languageArray
}

func StatsOfContributor(repoName, owner string) int {
	
	stats, _, _ := client.GHClient.Repositories.ListContributorsStats(client.GHContext, owner, repoName)

	totalCommitsOfUser := 0
	for _, v := range stats {
		if strings.Contains(v.GetAuthor().GetURL(), owner) {
			totalCommitsOfUser += v.GetTotal()
		}
	}

	fmt.Println(totalCommitsOfUser)
	
	return totalCommitsOfUser
}

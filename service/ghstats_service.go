package service

import (
	"fmt"
	"github.com/google/go-github/github"
	"gitsee/client"
	"gitsee/models"
	"strings"
	"sync"
	"time"
)

var (
	DistinctLanguagesofUser []string
)

func GetRepoStats(owner string) []models.AbstractRepo {

	stats, res, err := client.GHClient.Repositories.List(client.GHContext, owner, &github.RepositoryListOptions{
		ListOptions: github.ListOptions{
			Page: 1,
		},
	})

	if err != nil {
		fmt.Printf("Problem in getting repository information %v\n", err)
		return nil
	}

	for res.NextPage <= res.LastPage {
		nextStats, _, err := client.GHClient.Repositories.List(client.GHContext, owner, &github.RepositoryListOptions{
			ListOptions: github.ListOptions{
				Page: res.NextPage,
			},
		})

		if err != nil {
			fmt.Printf("Problem in getting repository information %v\n", err)
			return nil
		}

		for _, r := range nextStats {
			stats = append(stats, r)
		}

		res.NextPage += 1
	}

	repoStats := make([]models.AbstractRepo, 0)

	var wg sync.WaitGroup
	
	fmt.Println(len(stats))
	wg.Add(len(stats))

	GetRepoStats := func() {
		for _, v := range stats {
			go func(v *github.Repository) {
				defer wg.Done()
				// languageStats, _, _ := client.GHClient.Repositories.ListLanguages(client.GHContext, owner, v.GetName())
				
				start := time.Now()
				time.Sleep(time.Millisecond*50)
				userCommitCount := StatsOfContributor(v.GetName(), owner)
				fmt.Println(time.Since(start))
				
				repoStat := &models.AbstractRepo{
					RepoName:   v.GetName(),
				//	Languages:  languageStats,
					StarCount:  v.GetStargazersCount(),
					ForksCount: v.GetForksCount(),
					UserCommitCount: userCommitCount,
				}
				
				repoStats = append(repoStats, *repoStat)
			}(v)
		}

		wg.Wait()
	}

	GetRepoStats()

	// GetDistinctLanguagesOfUsersRepos(repoStats)

	return repoStats
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

	stats, _, err := client.GHClient.Repositories.ListContributorsStats(client.GHContext, owner, repoName)

	if err != nil {
		fmt.Printf("Problem in getting repository information %v\n", err)
		return 0
	}

	totalCommitsOfUser := 0
	for _, v := range stats {
		if strings.Contains(v.GetAuthor().GetURL(), owner) {
			totalCommitsOfUser = v.GetTotal()
		}
	}

	return totalCommitsOfUser
}

func ReposForks(repoStats []models.AbstractRepo) map[string]interface{} {
	reposForks := make(map[string]interface{}, 0)

	for _, v := range repoStats {
		if v.ForksCount > 0 {
			reposForks[v.RepoName] = v.ForksCount
		}
	}

	return reposForks
}

func RepoStars(repoStats []models.AbstractRepo) map[string]interface{} {
	reposStars := make(map[string]interface{}, 0)
	
	for _, v := range repoStats {
		if v.StarCount > 0 {
			reposStars[v.RepoName] = v.StarCount
		}
	}
	
	return reposStars
}

func RepoCommits(repoStats []models.AbstractRepo) map[string]interface{} {
	reposCommits := make(map[string]interface{}, 0)
	
	for _, v := range repoStats {
		if v.UserCommitCount > 0 {
			reposCommits[v.RepoName] = v.UserCommitCount
		}
	}
	
	return reposCommits
}

func FrequencyOfLanguages(repoStats []models.AbstractRepo) map[string]int {
	languageFreqs := make(map[string]int)

	for _, language := range DistinctLanguagesofUser {
		for _, repo := range repoStats {
			for languageField, _ := range repo.Languages {
				if language == languageField {
					languageFreqs[language] += 1
				}
			}
		}
	}

	return languageFreqs
}

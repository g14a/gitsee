package service

import (
	json2 "encoding/json"
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/google/go-github/github"
	"gitsee/client"
	"gitsee/models"
	"strings"
	"sync"
)

var (
	ReposOwnedByUser        []string
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
	wg.Add(len(stats))

	GetRepoStats := func() {
		for _, v := range stats {
			v := v
			go func() {
				defer wg.Done()
				languageStats, _, _ := client.GHClient.Repositories.ListLanguages(client.GHContext, owner, v.GetName())

				repoStat := &models.AbstractRepo{
					RepoName:   v.GetName(),
					Languages:  languageStats,
					StarCount:  v.GetStargazersCount(),
					ForksCount: v.GetForksCount(),
				}

				repoStat.UserCommitCount = StatsOfContributor(v.GetName(), owner)
				repoStats = append(repoStats, *repoStat)
			}()
		}

		wg.Wait()
	}

	GetRepoStats()

	GetDistinctLanguagesOfUsersRepos(repoStats)

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
			totalCommitsOfUser += v.GetTotal()
		}
	}

	return totalCommitsOfUser
}

// done
func ReposStarsAndForks(repoStats []models.AbstractRepo) map[string]interface{} {
	reposStars := make(map[string]interface{}, 0)

	for _, v := range repoStats {
		reposStars[v.RepoName] = map[string]interface{}{
			"stars":   v.StarCount,
			"forks":   v.ForksCount,
			"commits": v.UserCommitCount,
		}
	}

	return reposStars
}

func FrequencyOfLanguages(repoStats []models.AbstractRepo) {
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

	json, _ := json2.Marshal(languageFreqs)
	fmt.Println(string(json))
}

func UserDetails(user string) (map[string]interface{}, error) {
	ghUser, _, err := client.GHClient.Users.Get(client.GHContext, user)

	if err != nil {
		fmt.Printf("Problem in getting user information %v\n", err)
		return nil, err
	}

	return map[string]interface{}{
		"user": map[string]interface{}{
			"name":      ghUser.GetName(),
			"joined":    "Joined GitHub " + humanize.Time(ghUser.GetCreatedAt().Time),
			"location":  ghUser.GetLocation(),
			"avatar":    ghUser.GetAvatarURL(),
			"url":       ghUser.GetHTMLURL(),
			"followers": ghUser.GetFollowers(),
		},
	}, nil
}

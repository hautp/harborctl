package repository

import (
	"encoding/json"
	"fmt"
	middle "harborctl/pkg/middleware"
	pj "harborctl/pkg/project"

	"github.com/spf13/viper"
)

func ListRepositories(url string, project string) {
	var (
		hr       HarborRepository
		msg      string
		username string = viper.GetString("USERNAME")
		password string = viper.GetString("PASSWORD")
	)

	// Get total_page based on total_repo
	total_repo := pj.GetTotalRepo(url, project)
	msg = fmt.Sprintf("[Harbor] Found %d repositories:\n", total_repo)
	total_page := (total_repo / 50) + 1
	for page := 1; page <= total_page; page++ {
		hr_url := fmt.Sprintf("%s/api/v2.0/projects/%s/repositories?page=%d&page_size=50", url, project, page)
		repositories := middle.GetAPI(hr_url, username, password)
		json.Unmarshal([]byte(repositories), &hr)
		for _, r := range hr {
			msg += fmt.Sprintf("- %s (%d artifact)\n",
				r.Name,
				r.ArtifactCount)
		}
	}
	fmt.Println(msg)
}

func GetRepositories(url string, project string) []string {
	var (
		hr       HarborRepository
		rl       []string
		username string = viper.GetString("USERNAME")
		password string = viper.GetString("PASSWORD")
	)

	// Get total_page based on total_repo
	total_repo := pj.GetTotalRepo(url, project)
	total_page := (total_repo / 50) + 1
	for page := 1; page <= total_page; page++ {
		hr_url := fmt.Sprintf("%s/api/v2.0/projects/%s/repositories?page=%d&page_size=50", url, project, page)
		repositories := middle.GetAPI(hr_url, username, password)
		json.Unmarshal([]byte(repositories), &hr)
		for _, r := range hr {
			rl = append(rl, r.Name)
		}
	}
	return rl
}

package project

import (
	"encoding/json"
	"fmt"
	middle "harborctl/pkg/middleware"
	"log"

	"github.com/spf13/viper"
)

func ListProjects(url string) {
	var (
		hp       HarborProject
		msg      string
		username string = viper.GetString("USERNAME")
		password string = viper.GetString("PASSWORD")
		pj_url   string = fmt.Sprintf("%s/api/v2.0/projects?page=1&page_size=10", url)
	)

	projects := middle.GetAPI(pj_url, username, password)
	json.Unmarshal([]byte(projects), &hp)
	msg = fmt.Sprintf("[Harbor] List projects:\n")
	for n, p := range hp {
		msg += fmt.Sprintf("%d/ Project: %s (id: %d)\nOwner: %s (id: %d)\nPublic: %s\nTotal: %d repositories\n\n",
			n+1,
			p.Name,
			p.ProjectID,
			p.OwnerName,
			p.OwnerID,
			p.Metadata["public"],
			p.RepoCount)
	}
	fmt.Println(msg)
}

func GetProject(url string, project string) {
	var (
		hp       HarborProject
		msg      string
		username string = viper.GetString("USERNAME")
		password string = viper.GetString("PASSWORD")
		pj_url   string = fmt.Sprintf("%s/api/v2.0/projects?name=%s", url, project)
	)

	projects := middle.GetAPI(pj_url, username, password)
	json.Unmarshal([]byte(projects), &hp)
	if len(hp) != 0 {
		msg = fmt.Sprintf("[Harbor] Get project:\n")
		for _, p := range hp {
			msg += fmt.Sprintf("- Project: %s (id: %d)\nOwner: %s (id: %d)\nPublic: %s\nTotal: %d repositories\n",
				p.Name,
				p.ProjectID,
				p.OwnerName,
				p.OwnerID,
				p.Metadata["public"],
				p.RepoCount)
		}
	}
	fmt.Println(msg)
}

func GetTotalRepo(url string, project string) (total int) {
	var (
		hp       HarborProject
		pj_url   string = fmt.Sprintf("%s/api/v2.0/projects?name=%s", url, project)
		username string = viper.GetString("USERNAME")
		password string = viper.GetString("PASSWORD")
	)

	projects := middle.GetAPI(pj_url, username, password)
	json.Unmarshal([]byte(projects), &hp)

	if len(hp) == 0 {
		log.Fatalln("Could not get total repository.")
	} else {
		for i := range hp {
			total = hp[i].RepoCount
		}
	}
	return total
}

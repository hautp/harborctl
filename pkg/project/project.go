package project

import (
	"encoding/json"
	"fmt"
	middle "harborctl/pkg/middleware"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/viper"
)

func ListProjects(url string) {
	var (
		hp       HarborProject
		username string = viper.GetString("USERNAME")
		password string = viper.GetString("PASSWORD")
		pj_url   string = fmt.Sprintf("%s/api/v2.0/projects?page=1&page_size=10", url)
	)

	projects := middle.GetAPI(pj_url, username, password)
	json.Unmarshal([]byte(projects), &hp)
	if len(hp) != 0 {
		t := table.NewWriter()
		rowConfigAutoMerge := table.RowConfig{AutoMerge: true}
		t.AppendHeader(table.Row{"Project Name", "Owner", "Public", "Total repository"})
		for _, p := range hp {
			t.AppendRow(
				table.Row{
					fmt.Sprintf("%s (id: %d)", p.Name, p.ProjectID),
					fmt.Sprintf("%s (id: %d)", p.OwnerName, p.OwnerID),
					p.Metadata["public"],
					p.RepoCount},
				rowConfigAutoMerge,
			)
		}
		t.SetColumnConfigs([]table.ColumnConfig{{Number: 1, AutoMerge: true}})
		t.SetAutoIndex(true)
		t.SetStyle(table.StyleLight)
		t.SetOutputMirror(os.Stdout)
		t.Style().Options.SeparateRows = true
		t.Render()
	}
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
		fmt.Printf("[!] Something went wrong. Could not get repository of project \"%[1]s\". Please check the project name again.\n", project)
		os.Exit(1)
	} else {
		for i := range hp {
			total = hp[i].RepoCount
		}
	}
	return total
}

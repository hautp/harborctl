package repository

import (
	"encoding/json"
	"fmt"
	middle "harborctl/pkg/middleware"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/viper"
)

// List all artifacts of repository to file
func ListAllArtifactsToFiles(url string, project string) {
	rps := GetRepositories(url, project)
	for _, rp := range rps {
		fmt.Printf("- Repo: %s\nTags: %s\n\n", rp, getTagsRepository(url, rp))
	}
}

// List artifacts of specific repository as table format
func ListArtifactsToTable(url string, repository string) {
	detailsArtifact(url, repository)
}

func getTagsRepository(url string, repository string) []string {
	var (
		ha       HarborArtifact
		arts     []string
		username string = viper.GetString("USERNAME")
		password string = viper.GetString("PASSWORD")
	)

	r := strings.Split(repository, "/")
	prj, rp := r[0], r[1]
	ha_url := fmt.Sprintf("%s/api/v2.0/projects/%s/repositories/%s/artifacts?page=1&page_size=50&with_tag=true&with_label=false&with_scan_overview=false&with_signature=false&with_immutable_status=false", url, prj, rp)
	artifacts := middle.GetAPI(ha_url, username, password)
	json.Unmarshal([]byte(artifacts), &ha)
	for _, a := range ha {
		for n := range a.Tags {
			arts = append(arts, a.Tags[n].Name)
		}
	}
	// Sort arts slices
	sort.Sort(sort.Reverse(sort.StringSlice(arts)))
	sort.Slice(arts, func(i, j int) bool {
		if arts[i][:2] != arts[j][:2] {
			return arts[i] < arts[j]
		}
		artA, _ := strconv.Atoi(arts[i][:2])
		artB, _ := strconv.Atoi(arts[j][:2])
		return artA < artB
	})
	return arts
}

func detailsArtifact(url string, repository string) {
	arts := getTagsRepository(url, repository)
	fmt.Printf("[Harbor] List artifacts of \"%s\" repository:\n", repository)

	t := table.NewWriter()
	rowConfigAutoMerge := table.RowConfig{AutoMerge: true}
	t.AppendHeader(table.Row{"Repository", "Tags"})
	for _, art := range arts {
		t.AppendRow(table.Row{repository, art}, rowConfigAutoMerge)
	}

	t.SetColumnConfigs([]table.ColumnConfig{{Number: 1, AutoMerge: true}})
	t.SetStyle(table.StyleLight)
	t.SetOutputMirror(os.Stdout)
	t.Style().Options.SeparateRows = true
	t.Render()
}

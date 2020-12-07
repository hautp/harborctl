package healthcheck

import (
	"encoding/json"
	"fmt"
	middle "harborctl/pkg/middleware"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/viper"
)

func HealthCheck(url string) {
	var (
		hc       HarborSVC
		username string = viper.GetString("USERNAME")
		password string = viper.GetString("PASSWORD")
		hc_url   string = fmt.Sprintf("%s/api/v2.0/health", url)
	)

	resp := middle.GetAPI(hc_url, username, password)
	json.Unmarshal([]byte(resp), &hc)
	components := hc.Components
	if len(components) != 0 {
		t := table.NewWriter()
		rowConfigAutoMerge := table.RowConfig{AutoMerge: true}
		t.AppendHeader(table.Row{"Service", "Status"})
		for i := range components {
			t.AppendRow(table.Row{components[i].Name, components[i].Status}, rowConfigAutoMerge)
		}
		t.SetColumnConfigs([]table.ColumnConfig{{Number: 1, AutoMerge: true}})
		t.SetStyle(table.StyleLight)
		t.SetAutoIndex(true)
		t.SetOutputMirror(os.Stdout)
		t.Style().Options.SeparateRows = true
		t.Render()
	}
}

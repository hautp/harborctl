package users

import (
        "encoding/json"
        "fmt"
        middle "harborctl/pkg/middleware"
        "os"

        "github.com/jedib0t/go-pretty/v6/table"
        "github.com/spf13/viper"
)

func GetAllUsers (url string) {
    //get all users: https://hub.infra.tiki.services/api/v2.0/users
    var (
        userInfo    UserInfo
        username    string = viper.GetString("USERNAME")
        password    string = viper.GetString("PASSWORD")
        pj_url      string = fmt.Sprintf("%s/api/v2.0/users", url) 
    )

    projects := middle.GetAPI(pj_url, username, password)
    json.Unmarshal([]byte(projects), &userInfo)
    
    if len(userInfo) == 0 {
        fmt.Printf("[!] Something went wrong. Don't have any user in your registry\n")
    } else {
	t := table.NewWriter()
	rowConfigAutoMerge := table.RowConfig{AutoMerge: true}
	t.AppendHeader(table.Row{"User Name", "Email", "Admin Permission"})
	for _, u := range userInfo {
	    t.AppendRow(
	        table.Row{
	    	fmt.Sprintf("%s", u.UserName),
	    	fmt.Sprintf("%s", u.Email),
	    	fmt.Sprintf("%t", u.SysadFlag)},
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

func GetUser (url string, user string) {
    //filter username: https://hub.infra.tiki.services/api/v2.0/users?username=phat.nguyen2
    var (
        userInfo    UserInfo
        username    string = viper.GetString("USERNAME")
        password    string = viper.GetString("PASSWORD")
        pj_url      string = fmt.Sprintf("%s/api/v2.0/users?username=%s", url, user)
    )
    projects := middle.GetAPI(pj_url, username, password)
    json.Unmarshal([]byte(projects), &userInfo)

    if len(userInfo) == 0 {
        fmt.Printf("[!] Something went wrong. Don't have user \"%s\" in your registry\n", user)
    } else {
        t := table.NewWriter()
        rowConfigAutoMerge := table.RowConfig{AutoMerge: true}
        t.AppendHeader(table.Row{"User Name", "Email", "Admin Permission"})
        for _, u := range userInfo {
            t.AppendRow(
                table.Row{
                fmt.Sprintf("%s", u.UserName),
                fmt.Sprintf("%s", u.Email),
                fmt.Sprintf("%t", u.SysadFlag)},
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

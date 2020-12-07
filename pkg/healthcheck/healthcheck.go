package healthcheck

import (
	"encoding/json"
	"fmt"
	middle "harborctl/pkg/middleware"

	"github.com/spf13/viper"
)

func HealthCheck(url string) {
	var (
		msg      string
		hc       HarborSVC
		username string = viper.GetString("USERNAME")
		password string = viper.GetString("PASSWORD")
		hc_url   string = fmt.Sprintf("%s/api/v2.0/health", url)
	)

	resp := middle.GetAPI(hc_url, username, password)
	json.Unmarshal([]byte(resp), &hc)
	msg = fmt.Sprintf("[Harbor] Components status:\n")
	components := hc.Components
	for i := range components {
		if components[i].Status == "healthy" {
			msg += fmt.Sprintf("[+] Service: %s - Status: %s\n",
				components[i].Name, components[i].Status)
		} else {
			msg += fmt.Sprintf("[!] Service: %s | Status: %s\n",
				components[i].Name, components[i].Status)
		}
	}
	fmt.Println(msg)
}

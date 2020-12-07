package middleware

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func GetAPI(url string, username string, password string) []byte {
	var body []byte
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.SetBasicAuth(username, password)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	statusCode := int(resp.StatusCode)
	if statusCode == 200 {
		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()
	} else {
		fmt.Printf("Error: %s\n", resp.Status)
		defer resp.Body.Close()
		os.Exit(1)
	}
	return body
}

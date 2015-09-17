package main

import (
	"encoding/json"
	"fmt"
	"github.com/DigitalInnovation/bilious-funicular/Godeps/_workspace/src/github.com/Bowery/slack"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/DigitalInnovation/bilious-funicular/global"
)

var (
	client *slack.Client
)

type test_struct struct {
	Action        string `json:"action"`
	Number        int    `json:"number"`
	Pull_requests struct {
		Html_url string `json:"html_url"`
	} `json:"pull_request"`
	Repositories struct {
		Name string `json:"name"`
	} `json:"repository"`
	Sender struct {
		Login string `json:"login"`
	} `json:"sender"`
}

func test(rw http.ResponseWriter, req *http.Request) {

	decoder := json.NewDecoder(req.Body)
	for {
		var t test_struct
		if err := decoder.Decode(&t); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		log.Printf("%s\n", "Working on it")

		client = slack.NewClient(global.Config.Slack_Key)
		if t.Action == "opened" {
			err := client.SendMessage("#staff-ass-apps", "PR please \n"+t.Repositories.Name+" "+t.Pull_requests.Html_url, t.Sender.Login)
			if err != nil {
				log.Fatal(err)
			}
		} else if t.Action == "closed" {
			err := client.SendMessage("#staff-ass-apps", t.Repositories.Name+"\n"+"PR is closed", t.Sender.Login)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			err := client.SendMessage("#testhooks", t.Repositories.Name+"\n"+"Something wierd is going on", t.Sender.Login)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
func main() {
	logEnvironmentVariables()

	global.Setup()

	http.HandleFunc("/", test)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", global.Config.Port), nil))
}

func logEnvironmentVariables() {
	log.Printf("PORT: %v", os.Getenv("PORT"))
	log.Printf("SLACKKEY: %v", os.Getenv("SLACKKEY"))
}

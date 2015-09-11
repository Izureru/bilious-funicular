package main

import (
	"encoding/json"
	"github.com/Bowery/slack"
	"io"
	"log"
	"net/http"
)

var (
	client *slack.Client
)

type test_struct struct {
	Action        string `json:"action"`
	Number        int    `json:"number"`
	Pull_requests struct {
		Url string `json:"url"`
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

		client = slack.NewClient("xxxxxxx")
		err := client.SendMessage("#staff-ass-apps", t.Repositories.Name+" "+t.Pull_requests.Url, t.Sender.Login)
		if err != nil {
			log.Fatal(err)
		}
	}
}
func main() {

	http.HandleFunc("/", test)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

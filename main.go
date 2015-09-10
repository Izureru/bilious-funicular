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
		Diff_url string `json:"diff_url"`
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

		client = slack.NewClient("xxxxxxxxx")
		err := client.SendMessage("#testhooks", t.Repositories.Name+" "+t.Pull_requests.Diff_url, t.Sender.Login)
		if err != nil {
			log.Fatal(err)
		}
	}
}
func main() {

	http.HandleFunc("/", test)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

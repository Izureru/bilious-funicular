package global

import (
	"os"
	"strconv"
)

var (
	Config ConfigStruct
)

type ConfigStruct struct {
	Port      int
	Slack_Key string
}

func loadConfig() {

	Config = ConfigStruct{}

	port, _ := strconv.Atoi(os.Getenv("PORT"))

	Config.Port = port
	Config.Slack_Key = os.Getenv("SLACKKEY")
}

func Setup() {
	loadConfig()
}

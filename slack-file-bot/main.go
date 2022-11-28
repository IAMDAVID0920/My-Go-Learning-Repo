package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	bot_token := ""
	channel_id := ""
	file_path := "resume.pdf"
	os.Setenv("SLACK_BOT_TOKEN", bot_token)
	os.Setenv("CHANNLE_ID", channel_id)
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{file_path}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}

		fmt.Printf("Name: %s, URL: %s\n", file.Name, file.URL)
	}
}

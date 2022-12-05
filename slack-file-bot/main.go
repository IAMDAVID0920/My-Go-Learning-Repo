package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	bot_token := "xoxb-4426694404677-4429658202418-J6fOD9m1pqVE5zthttjXMtic"
	channel_id := "C04CMJUFNBU" // # general
	// file_path := "resume.pdf"
	os.Setenv("SLACK_BOT_TOKEN", bot_token)
	os.Setenv("CHANNLE_ID", channel_id)
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"a.txt"}

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

	for i := 0; i < 20; i++ {
		fmt.Printf("i: %d", i)
	}
}

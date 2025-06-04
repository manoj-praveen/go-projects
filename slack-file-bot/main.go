package main

import (
	"fmt"
	"github.com/slack-go/slack"
	"os"
)

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "test_token")
	os.Setenv("CHANNAL_ID", "")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channel := os.Getenv("CHANNEL_ID")
	fileArr := []string{"dummy.txt"}

	for i := range fileArr {
		println(fileArr[i])
		params := slack.UploadFileV2Parameters{
			Channel: channel,
			Filename:    fileArr[i],
			FileSize: 1,
		}
		file, err := api.UploadFileV2(params)
		if err != nil {
			fmt.Printf("Error in uploading file to slack %s\n", err)
			return
		}
		fmt.Printf("File ID: %s, Title: %s\n", file.ID, file.Title)
	}
}

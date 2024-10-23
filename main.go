package main

import (
        "fmt"
        "os"

        "github.com/slack-go/slack"
)

func main() {
        // Set the Slack Bot Token and Channel ID
        os.Setenv("SLACK_BOT_TOKEN", "") // Put your SLACK BOT TOKEN here
        os.Setenv("CHANNEL_ID", "")      // Put your CHANNEL ID here

        // Create a new Slack API client using the token
        api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
        channelID := os.Getenv("CHANNEL_ID")

        // Array of file paths to upload
        fileArr := []string{"crime_data.csv"}

        // Loop through and upload each file
        for i := 0; i < len(fileArr); i++ {
                params := slack.FileUploadParameters{
                        Channels: []string{channelID}, // Correctly pass the channel ID in a slice
                        File:     fileArr[i],
                } // <-- Missing closing curly brace

                // Upload the file to the Slack channel
                file, err := api.UploadFile(params)
                if err != nil {
                        fmt.Printf("Error uploading file: %v\n", err)
                        continue // Skip to the next file if there's an error
                }

                // Print the name and URL of the uploaded file
                fmt.Printf("Name: %s, URL: %s\n", file.Name, file.URL)
        }
}
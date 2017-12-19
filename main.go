package main

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"./slack"
	"./storeReview"
)

const location = "Asia/Tokyo"

func main() {
	storeURL := "https://itunes.apple.com/jp/rss/customerreviews/id=" + os.Getenv("APPSTORE_ID") + "/xml"
	data := httpGet(storeURL)

	result := storeReview.XML{}
	err := xml.Unmarshal([]byte(data), &result)
	if err != nil {
		log.Fatal(err)
		return
	}

	baseTime := time.Now().Add(-time.Hour * 24)

	webhookUrl := os.Getenv("REVIEW_SLACK_WEBHOOK_URL")
	for _, review := range result.Reviews {
		updatedAt, err := time.Parse(
			"2006-01-02T15:04:05-07:00",
			review.Updated)
		if err != nil {
			log.Fatal(err)
		}

		if updatedAt.Before(baseTime) || int(review.Rating) == 0 {
			continue
		}

		attachment := slack.Attachment{}
		version := review.Version
		star := strings.Repeat("⭐", int(review.Rating))
		title := review.Title
		comment := review.Comment[0].Text
		author := review.Author.Name

		loc, timeErr := time.LoadLocation(location)
		if timeErr != nil {
			log.Fatal(timeErr)
		}
		updatedStr := updatedAt.In(loc).Format("2006/01/02 Mon 15:04:05")

		attachment.AddField(slack.Field { Title: star + " " + title, Value: comment })
		payload := slack.Payload {
			Text: "App Store Review\n" + "ver: " + version + " [" + updatedStr + "]",
			Username: author,
			Channel: os.Getenv("REVIEW_SLACK_CHANNEL"),
			IconEmoji: os.Getenv("REVIEW_SLACK_ICON"),
			Attachments: []slack.Attachment{attachment},
		}
		slackErr := slack.Send(webhookUrl, payload)
		if len(slackErr) > 0 {
			log.Fatal(slackErr)
		}
	}

}

func httpGet(url string) string {
	response, _ := http.Get(url)
	body, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	return string(body)
}

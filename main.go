package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"

	"./storeReview"
)

func main() {
	storeURL := "https://itunes.apple.com/jp/rss/customerreviews/id=" + os.Getenv("APPSTORE_ID") + "/xml"
	data := httpGet(storeURL)

	result := storeReview.XML{}
	err := xml.Unmarshal([]byte(data), &result)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	for _, review := range result.Reviews {
		fmt.Printf("%v\n", review.Id)
		fmt.Printf("%v\n", review.Title)
		fmt.Printf("%v\n", review.Updated)
		fmt.Printf("%v\n", review.Comment[0].Text)
		fmt.Printf("%d\n", review.Rating)
		fmt.Printf("%v\n", review.Version)
		fmt.Printf("%v\n", review.Author.Name)
		fmt.Println()
	}
}

func httpGet(url string) string {
	response, _ := http.Get(url)
	body, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	return string(body)
}

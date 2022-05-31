package main

import "fmt"
import "net/http"
import "encoding/json"
import "log"

const BaseURL = "https://hacker-news.firebaseio.com/v0/"
const TopStoriesURL = BaseURL + "topstories.json"
const ItemURL = BaseURL + "item/%d.json"
const MinimumScore = 150
const MaxArticles = 30

type Item struct {
	By          string `json:"by"`
	Descendants int    `json:"descendants"`
	Id          int    `json:"id"`
	Kids        []int  `json:"kids"`
	Score       int    `json:"score"`
	Time        int    `json:"time"`
	Title       string `json:"title"`
	Item_type   string `json:"type"`
	Url         string `json:"url"`
}

func getTopStories() {
	resp, err := http.Get(TopStoriesURL)
	if err != nil {
		log.Fatal(err)
	}

	decoder := json.NewDecoder(resp.Body)
	var v []int

	err = decoder.Decode(&v)

	for _, v := range v[:MaxArticles] {
		var item Item
		url := fmt.Sprintf(ItemURL, v)
		http_response, _ := http.Get(url)

		decoder := json.NewDecoder(http_response.Body)

		decoder_error := decoder.Decode(&item)

		if decoder_error != nil {
			log.Fatal(decoder_error)
		}

		// Only print stories above minimum score
		if item.Score > MinimumScore && item.Item_type == "story" {
			printStory(item.Title, item.Url)
		}
	}
}

func printStory(Title string, Url string) {
	fmt.Printf("%s - %s\n", Title, Url)
}

func main() {
	getTopStories()
}

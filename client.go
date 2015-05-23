package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var BaseUrl = "https://api.twitter.com"

type Status struct {
	Text      string `json:"text"`
	Retweets  int    `json:"retweet_count"`
	Favorites int    `json:"favorite_count"`
	User      struct {
		Name       string `json:"name"`
		ScreenName string `json:"screen_name"`
	} `json:"user"`
	Retweet struct {
		User struct {
			Name       string `json:"name"`
			ScreenName string `json:"screen_name"`
		} `json:"user"`
	} `json:"retweeted_status"`
	Entities struct {
		Tags []struct {
			Text string `json:"text"`
		} `json:"hashtags"`
	} `json:"entities"`
}

func Search(query, token string) (statuses []Status, err error) {
	body, err := fetchQuery(query, token)
	if err != nil {
		return
	}

	var tweets struct {
		Statuses []Status `json:"statuses"`
	}

	err = json.Unmarshal(body, &tweets)
	if err != nil {
		return
	}

	statuses = tweets.Statuses
	return
}

func fetchQuery(query, token string) (body []byte, err error) {
	q := url.Values{}
	q.Set("count", "5")
	q.Set("q", query)

	searchUrl := BaseUrl + "/1.1/search/tweets.json?" + q.Encode()
	_, body, err = fetchUrl(searchUrl, token)
	return
}

func fetchUrl(url, token string) (res *http.Response, body []byte, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	req.Header.Set("Authorization", "Bearer "+token)
	res, body, err = fetch(req)
	return
}

func fetch(req *http.Request) (res *http.Response, body []byte, err error) {
	log.Println("client:", req.URL)

	req.Header.Set("User-Agent", "whatever/challenge")
	res, err = http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	defer res.Body.Close()
	log.Println("client:", res.Status)

	body, err = ioutil.ReadAll(res.Body)
	return
}

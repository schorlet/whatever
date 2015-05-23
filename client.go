package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

var BaseUrl = "https://api.twitter.com"

func main() {
	key := os.Getenv("CONSUMER_KEY")
	secret := os.Getenv("CONSUMER_SECRET")

	if key == "" || secret == "" {
		log.Fatal("client: missing consumer credentials")
	}

	token, err := GetToken(key, secret)
	if err != nil {
		log.Fatal(err)
	}

	body, err := Search("#golang", token)
	if err != nil {
		log.Fatal(err)
	}

	var out bytes.Buffer
	json.Indent(&out, body, "", "\t")
	out.WriteTo(os.Stdout)
}

func Search(query, token string) (body []byte, err error) {
	q := url.Values{}
	q.Set("count", "2")
	q.Set("result_type", "popular")
	q.Set("include_entities", "false")
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

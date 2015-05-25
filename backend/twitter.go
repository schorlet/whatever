package backend

import (
	"fmt"
	"strings"
	"time"
)

var BaseUrl = "https://api.twitter.com"

type Status struct {
	Text      string    `json:"text"`
	Created   CreatedAt `json:"created_at"`
	Retweets  int       `json:"retweet_count"`
	Favorites int       `json:"favorite_count"`
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

type CreatedAt time.Time

func (t CreatedAt) Layout() string {
	return "Mon Jan 02 15:04:05 -0700 2006"
}

func (t CreatedAt) MarshalJSON() ([]byte, error) {
	str := time.Time(t).Format(t.Layout())
	str = fmt.Sprintf(`"%s"`, str)
	return []byte(str), nil
}

func (t *CreatedAt) UnmarshalJSON(b []byte) error {
	str := strings.Trim(string(b), `"`)
	date, err := time.Parse(t.Layout(), str)
	if err != nil {
		return err
	}
	*t = CreatedAt(date)
	return nil
}

package twitter

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type users struct {
}

type follower struct {
	users               []users
	next_cursor         int
	next_cursor_str     string
	previous_cursor     int
	previous_cursor_str string
}

type Twitter struct {
	AccountName string
	follower    []string
	following   []string
	client      *http.Client
}

func (t Twitter) GetRelation() {
	t.client = &http.Client{}
	t.getFollower()
}

func (t Twitter) getFollower() {
	req, err := http.NewRequest("GET", os.Getenv("TWITTER_API")+"followers/list.json?screen_name="+t.AccountName, nil)
	if err != nil {
		log.Fatal("getFollower: cannot build packet")
		return
	}
	req.Header.Set("authorization", "Bearer "+os.Getenv("BEARER_TOKEN"))
	resp, err := t.client.Do(req)
	if err != nil {
		log.Fatal("getFollower: cannot get twitter: ")
		return
	}

	defer resp.Body.Close()
	var target interface{}
	json.NewDecoder(resp.Body).Decode(target)

}

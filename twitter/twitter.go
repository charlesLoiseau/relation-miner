package twitter

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"relation-miner/model"
	"relation-miner/neo4j"
	"strconv"
)

type Twitter struct {
	AccountName string
	Follower    []model.SimpleUser
	Following   []model.SimpleUser
	User        model.SimpleUser
	client      *http.Client
	neo         neo4j.Neo4j
	tmp         []int
}

func (t *Twitter) GetRelation() {
	t.client = &http.Client{}
	t.neo = neo4j.Neo4j{}
	t.neo.Init()
	t.getTarget()
	t.getFollower(-1)
	t.getFollowing(-1)
	t.neo.Close()
}

func (t *Twitter) getTarget() {
	target, err := t.getUserInfo(0, t.AccountName)
	if err != nil {
		log.Fatal("getTarget: " + err.Error())
		return
	}
	t.User.Name = target.Name
	t.User.Id = target.Id
}

func (t *Twitter) getFollower(cursor int) {
	req, err := http.NewRequest("GET", os.Getenv("TWITTER_API")+"followers/ids.json?screen_name="+t.AccountName+"&cursor="+strconv.Itoa(cursor)+"&count=5000", nil)
	if err != nil {
		log.Fatal("getFollower: cannot build packet" + err.Error())
		return
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("authorization", "Bearer "+os.Getenv("BEARER_TOKEN"))
	resp, err := t.client.Do(req)
	if err != nil {
		log.Fatal("getFollower: cannot get twitter " + err.Error())
		return
	}
	defer resp.Body.Close()
	var target model.FollowerId
	err = json.NewDecoder(resp.Body).Decode(&target)
	if err != nil {
		log.Fatal("getFollower: Error while parsing json " + err.Error())
		return
	}
	for _, id := range target.Ids {
		exist, err := t.neo.VerifyExist(id)
		if err != nil {
			log.Fatal("getFollower: " + err.Error())
			return
		}
		if !exist {
			t.tmp = append(t.tmp, id)
		} else {
			t.Follower = append(t.Follower, model.SimpleUser{Id: id, Name: ""})
		}
	}
	if target.NextCursor > 0 {
		t.getFollower(target.NextCursor)
	}
	for i := 0; i < len(t.tmp); i += 100 {
		var j = i + 100
		if j > len(t.tmp) {
			j = len(t.tmp)
		}
		us, err := t.getUsersInfo(t.tmp[i:j])
		if err != nil {
			log.Fatal("getFollower: " + err.Error())
			return
		}
		t.Follower = append(t.Follower, us...)
	}
}

func (t *Twitter) getUsersInfo(ids []int) ([]model.SimpleUser, error) {
	var query = "user_id="
	var users []model.SimpleUser
	for x, id := range ids {
		query += strconv.Itoa(id)
		if x < len(ids)-1 {
			query += ","
		}
	}
	req, err := http.NewRequest("GET", os.Getenv("TWITTER_API")+"users/lookup.json?"+query, nil)
	if err != nil {
		log.Fatal("getUserInfo: cannot build packet" + err.Error())
		return []model.SimpleUser{}, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("authorization", "Bearer "+os.Getenv("BEARER_TOKEN"))
	resp, err := t.client.Do(req)
	if err != nil {
		log.Fatal("getUserInfo: cannot get twitter " + err.Error())
		return []model.SimpleUser{}, err
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatal("getUserInfo: Error" + strconv.Itoa(resp.StatusCode))
	}
	defer resp.Body.Close()

	var results []model.User
	err = json.NewDecoder(resp.Body).Decode(&results)
	if err != nil {
		log.Fatal("getUserInfo: Error while parsing json " + err.Error())
		return []model.SimpleUser{}, err
	}
	for _, user := range results {
		users = append(users, model.SimpleUser{Id: user.Id, Name: user.Name})
	}
	return users, nil
}

func (t *Twitter) getUserInfo(id int, name string) (model.SimpleUser, error) {
	var query string
	if name == "" {
		query = "user_id=" + strconv.Itoa(id)
	} else {
		query = "screen_name=" + name
	}
	req, err := http.NewRequest("GET", os.Getenv("TWITTER_API")+"users/show.json?"+query, nil)
	if err != nil {
		log.Fatal("getUserInfo: cannot build packet" + err.Error())
		return model.SimpleUser{}, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("authorization", "Bearer "+os.Getenv("BEARER_TOKEN"))
	resp, err := t.client.Do(req)
	if err != nil {
		log.Fatal("getUserInfo: cannot get twitter " + err.Error())
		return model.SimpleUser{}, err
	}
	defer resp.Body.Close()
	var target model.User
	err = json.NewDecoder(resp.Body).Decode(&target)
	if err != nil {
		log.Fatal("getUserInfo: Error while parsing json " + err.Error())
		return model.SimpleUser{}, err
	}
	return model.SimpleUser{Id: target.Id, Name: target.Name}, nil
}

func (t *Twitter) getFollowing(cursor int) {
	req, err := http.NewRequest("GET", os.Getenv("TWITTER_API")+"friends/list.json?screen_name="+t.AccountName+"&cursor="+strconv.Itoa(cursor)+"&count=200", nil)
	if err != nil {
		log.Fatal("getFollower: cannot build packet" + err.Error())
		return
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("authorization", "Bearer "+os.Getenv("BEARER_TOKEN"))
	resp, err := t.client.Do(req)
	if err != nil {
		log.Fatal("getFollower: cannot get twitter " + err.Error())
		return
	}
	defer resp.Body.Close()
	var target model.Follower
	err = json.NewDecoder(resp.Body).Decode(&target)
	if err != nil {
		log.Fatal("getFollower: Error while parsing json " + err.Error())
		return
	}
	for _, u := range target.Users {
		t.Following = append(t.Following, model.SimpleUser{Id: u.Id, Name: u.Name})
	}
	if target.NextCursor > 0 {
		t.getFollowing(target.NextCursor)
	}
}

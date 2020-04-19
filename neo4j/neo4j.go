package neo4j

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"log"
	"os"
	"relation-miner/model"
)

type Neo4j struct {
	driver    neo4j.Driver
	session   neo4j.Session
	result    neo4j.Result
	err       error
	User      model.SimpleUser
	Following []model.SimpleUser
	Follower  []model.SimpleUser
}

func (n *Neo4j) Close() {
	defer n.driver.Close()
	defer n.session.Close()
}

func (n *Neo4j) Init() {
	n.driver, n.err = neo4j.NewDriver(os.Getenv("NEO4J"), neo4j.BasicAuth(os.Getenv("NEO4J_USER"), os.Getenv("NEO4J_PASSWORD"), ""), func(c *neo4j.Config) {
		c.Encrypted = false
	})
	if n.err != nil {
		log.Fatal("CreateRelation: " + n.err.Error())
	}

	n.session, n.err = n.driver.Session(neo4j.AccessModeWrite)
	if n.err != nil {
		log.Fatal("CreateRelation: " + n.err.Error())
	}

}

func (n *Neo4j) CreateRelation() {
	exist, err := n.VerifyExist(n.User.Id)
	if err != nil {
		log.Fatal("createNode->verifyExist: ", err.Error())
		return
	}
	if !exist {
		result, err := n.session.Run(`CREATE (n:USER {name:$name, id: $id })`, map[string]interface{}{
			"name": n.User.Name,
			"id":   n.User.Id,
		})
		if err != nil {
			log.Fatal("CreateRelation: ", err.Error())
			return
		}
		if result.Err() != nil {
			log.Fatal("CreateRelation: ", result.Err().Error())
		}
	}
	n.createUsers(n.Follower)
	n.createUsers(n.Following)
	for _, user := range n.Following {
		n.createRelation(n.User, user, "Following")
	}
	for _, user := range n.Follower {
		n.createRelation(user, n.User, "Follower")
	}
}

func (n *Neo4j) createUsers(users []model.SimpleUser) {
	for _, user := range users {
		exist, err := n.VerifyExist(user.Id)
		if err != nil {
			log.Fatal("createNode->verifyExist: ", err.Error())
			return
		}
		if !exist {
			result, err := n.session.Run(`CREATE (n:USER {name:$name, id: $id }) `,
				map[string]interface{}{
					"UserId": n.User.Id,
					"name":   user.Name,
					"id":     user.Id,
				})
			if err != nil {
				log.Fatal("createNode: ", err.Error())
				return
			}
			if result.Err() != nil {
				log.Fatal("createNode: ", result.Err().Error())
			}
		}

	}
}

func (n *Neo4j) VerifyExist(id int) (bool, error) {
	result, err := n.session.Run(`MATCH (a:USER) WHERE a.id = $id RETURN a`,
		map[string]interface{}{
			"id": id,
		})
	if err != nil {
		log.Fatal("verifyExist: ", err.Error())
		return false, err
	}
	if result.Err() != nil {
		log.Fatal("verifyExist: ", result.Err().Error())
		return false, result.Err()
	}
	if result.Next() {
		return true, nil
	} else {
		return false, nil
	}
}

func (n *Neo4j) createRelation(from model.SimpleUser, to model.SimpleUser, label string) {
	result, err := n.session.Run(`MATCH (a:USER), (b:USER) WHERE a.id = $fromId AND b.id = $toId CREATE (a)-[:`+label+`]->(b)`,
		map[string]interface{}{
			"fromId": from.Id,
			"toId":   to.Id,
			"label":  label,
		})
	if err != nil {
		log.Fatal("CreateRelation: ", err.Error())
		return
	}
	if result.Err() != nil {
		log.Fatal("CreateRelation: ", result.Err().Error())
	}
}

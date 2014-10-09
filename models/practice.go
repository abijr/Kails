package models

import (
	"errors"

	"bitbucket.com/abijr/kails/db"
	"github.com/diegogub/aranGO"
	"gopkg.in/mgo.v2/bson"
)

var (
	topics = db.Collection("topics")
	privileges = db.Collection("privileges")

	errLevelNotExists = errors.New("level does not exist")
)


//Structure that represents the topics collection in database
type Topic struct {
	aranGO.Document
	Id			int			`json:"Id"`
	Name		string		`json:"Name"`
	Subtopics	[]string	`json:"Subtopics"`
	NoUsers		int 		`json:"NoUsers"`
}

type Privilege struct {
	aranGO.Document
	Level		int 		`json:"Level"`
	Topics 		int 		`json:"Topics"`
	Friends		int 		`json:"Friends"`
	Features	[]string	`json:"Features"`
	Time		int 		`json:"Time"`
}

//Then, you search for the subtopics and the number of users 
//corresponding to the main topic retrieved previously
func GetTopicInfo(TopicToSearch string) ([]string, int){
	var topic *Topic
	topic = new(Topic)

	topics.First(bson.M{"Topic": TopicToSearch}, topic)

	return topic.Subtopics, topic.NoUsers
}

//Get the privileges of the user based on the level they are in
func GetUserPrivileges(level int) (*Privilege, error) {
	var privilege *Privilege
	privilege = new(Privilege)

	if level <= 0 {
		return nil, errLevelNotExists
	}

	privileges.First(bson.M{"Level": level}, privilege)

	return privilege, nil
}


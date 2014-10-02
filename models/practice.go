package models

import (
	"bitbucket.com/abijr/kails/db"
	"github.com/diegogub/aranGO"
	"gopkg.in/mgo.v2/bson"
)

var (
	topics = db.Collection("topics")
)


//Structure that represents the topics collection in database
type Topic struct {
	aranGO.Document
	Id			int			`json:"Id"`
	Name		string		`json:"Name"`
	Subtopics	[]string	`json:"Subtopics"`
	NoUsers		int 		`json:"NoUsers"`
}

//First, is necesary to get the information from the user
func GetUserInfo(Username string) *User {
	var user *User
	user = new(User)

	users.First(bson.M{"Username": Username}, user)

	return user
}

//Then, you search for the subtopics and the number of users 
//corresponding to the main topic retrieved previously
func GetTopicInfo(TopicToSearch string) ([]string, int){
	var topic *Topic
	topic = new(Topic)

	topics.First(bson.M{"Topic": TopicToSearch}, topic)

	return topic.Subtopics, topic.NoUsers
}


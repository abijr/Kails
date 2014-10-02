import (
	"errors"
	"fmt"


	"encoding/base64"
	"log"

	"bitbucket.com/abijr/kails/db"
	"bitbucket.com/abijr/kails/util"
	"github.com/diegogub/aranGO"
	"gopkg.in/mgo.v2/bson"
)

var (
	users = db.Collection("users")
	topics = db.Collection("topics")
)

// UserLevel is the representation of the
// user progress in a level
type UserLevel struct {
	// Id is the level Id
	Id int `json:"Id"`
	// LastPracticed is the list practiced time
	LastPracticed time.Time `json"Last"`
}

// User is the user structure, it holds user information
type User struct {
	aranGO.Document
	Username      string               `json:"Username"`
	Email         string               `json:"Email"`
	Password      string               `json:"Password"`
	Salt          string               `json:"Salt"`
	Language      string               `json:"Language"`		//Checked if I need all the variables
	StudyLanguage string               `json:"StudyLanguage"`
	Created       time.Time            `json:"Created"`
	Levels        map[string]UserLevel `json:"Levels"`
	Topics 		  []string			   'json:"Topics"'
}

//Structure that represents the topics collection in database
type Topic struct {
	aranGO.Document
	Id			int			'json:"Id"'
	Name		string		'json:"Name"'
	Subtopics	[]string	'json:"Subtopics"'
	NoUsers		int 		'json:"NoUsers"'
}

//First, is necesary to get the information from the user
func GetUserInfo(Username string) *User {
	var user *User
	user = new(User)

	err := users.First(bson.M("Username": Username), user)

	return user
}

//Then, you search for the subtopics and the number of users 
//corresponding to the main topic retrieved previously
func GetTopicInfo(Topic string) ([]string, int){
	var topic *Topic
	topic = new(Topic)

	err := topics.First(bson.M("Topic": Topic), topic)

	return topic.Subtopics, topic.NoUser
}


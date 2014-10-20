package models

import (
	"fmt"
	"log"

	"github.com/diegogub/aranGO"
	"gopkg.in/mgo.v2/bson"

	"bitbucket.com/abijr/kails/db"
)

const (
	// UserRelationsCollection is the name of the collection holding friendship/requests relations
	UserRelationsCollection = "relations"

	// UserRelationsGraph is the name of the graph holding the relations
	UserRelationsGraph = "relations"

	// Relationship types
	typeRequest    = "Request"
	typeFriendship = "Friendship"
)

var (
	// relations collections
	relations = db.Collection(UserRelationsCollection)

	// relations graph
	graph = db.DB.Graph(UserRelationsGraph)
)

type Relation struct {
	Id   string `json:"_id,omitempty"`
	Key  string `json:"_key,omitempty"`
	From string `json:"_from,omitempty"`
	To   string `json:"_to,omitempty"`
	Type string `json:"Type,omitempty"`
}

// SendFriendRequest envia una solicitud de amistad de
// parte de 'user' a 'other'.
func (user *User) SendFriendRequest(other *User) error {
	if user.Id == "" || other.Id == "" {
		return errRelationInvalid
	}

	relation := Relation{
		From: other.Id,
		To:   user.Id,
	}

	// Check if there's no current relationship
	err := relations.First(relation, nil)
	// If relation already exists
	log.Println("1st check: ", err)
	if err == nil {
		return errRelationInvalid
	}

	// The other way around
	relation.From, relation.To = relation.To, relation.From
	err = relations.First(relation, nil)
	// If relation already exists
	log.Println("2nd check", err)
	if err == nil {
		return errRelationInvalid
	}

	// Set the type of relationship
	relation.Type = typeRequest

	// Relation from `user` to `other`
	err = relations.SaveEdge(relation, user.Id, other.Id)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (user *User) AcceptFriendRequest(other *User) error {
	if user.Id == "" || other.Id == "" {
		return errRelationInvalid
	}

	log.Println(other.Id, user.Id)

	currentRelation := Relation{
		From: other.Id,
		To:   user.Id,
		Type: typeRequest,
	}

	var newRelation Relation

	// Fetch friend request, otherwise fail (no friendship if no request)
	err := relations.First(currentRelation, &newRelation)
	if err != nil {
		return errRelationInvalid
	}

	// Relation from user to other
	err = graph.PatchE(UserRelationsCollection, newRelation.Key, nil, bson.M{
		"Type": typeFriendship,
	})
	if err != nil {
		log.Println("Patch error: ", err)
		return err
	}

	return nil

}

// neighborQuery genera el query para encontrar
// a los vecinos del vertice que cumplen con la
// condicion "Type: {{relationType}}"
func neighborQuery(vertexId, relationType string) *aranGO.Query {

	q := fmt.Sprintf(
		"FOR e IN GRAPH_NEIGHBORS"+
			"('%v','%v', {direction: 'any', edgeExamples: [{Type: '%v'}]}) "+
			"RETURN {Username: e.vertex.Username, StudyLanguage: e.vertex.StudyLanguage}",
		UserRelationsGraph, // Graph name
		vertexId,           // Id of the user to get friends of
		relationType,       // Type of relationship we want
	)

	return aranGO.NewQuery(q)

}

type Friend struct {
	Username      string
	StudyLanguage string
}

// ListFriends regresa todos los amigos del usuario.
// De momento solo los imprime, es necesario establecer
// su correcto funcionamiento
func (user *User) ListFriends() ([]Friend, error) {

	q := neighborQuery(user.Id, typeFriendship)

	c, err := db.DB.Execute(q)
	if err != nil {
		log.Println("ListFriends error: ", err)
		return nil, err
	}

	var friendList []Friend
	c.FetchBatch(&friendList)

	log.Printf("Friends: %+v\n", friendList)
	return friendList, nil

}

func (user *User) ListRequests() {

	q := neighborQuery(user.Id, typeRequest)

	c, err := db.DB.Execute(q)
	if err != nil {
		log.Println("ListRequests error: ", err)
	}

	var friendList []string
	c.FetchBatch(&friendList)

	log.Println(friendList)

}

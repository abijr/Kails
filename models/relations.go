package models

import (
	"fmt"
	"log"

	"github.com/diegogub/aranGO"

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

func (user *User) SendFriendRequest(other *User) error {
	if user.Id == "" || other.Id == "" {
		return errRelationInvalid
	}

	relation := Relation{
		From: other.Id,
		To:   user.Id,
		Type: typeRequest,
	}

	// Check if there's no current relationship
	cur, _ := relations.Example(relation, 0, 1)
	// If relation already exists
	if cur.Count() != 0 {
		return errRelationInvalid
	}

	// Relation from `user` to `other`
	err := relations.SaveEdge(relation, user.Id, other.Id)
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

	err := relations.First(currentRelation, &newRelation)
	if err != nil {
		return errRelationInvalid
	}

	newRelation.Type = typeFriendship

	// Relation from user to other
	err = relations.Patch(newRelation.Id, newRelation.Type)
	if err != nil {
		log.Println("Patch error: ", err)
		return err
	}

	return nil

}

func (user *User) ListFriends() error {
	qString := fmt.Sprintf(
		"FOR c IN GRAPH_NEIGHBORS"+
			"('%v','%v', {direction: 'any', edgeExamples: [{Type: '%v'}]}) "+
			"RETURN e.Username",
		UserRelationsGraph, // Graph name
		user.Id,            // Id of the user to get friends of
		typeFriendship,     // Type of relationship we want
	)

	log.Println(qString)

	q := aranGO.NewQuery(qString)

	c, err := db.DB.Execute(q)

	return nil
}

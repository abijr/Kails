package models

import (
	"log"

	"gopkg.in/mgo.v2/bson"
)

// AddExperience adds experience points to user
// if the level threshold is reached, levels up
// user
func (user *User) AddExperience(xp int) error {
	newExperience := user.Experience + xp
	newLevel := user.Level
	if newExperience >= ExperienceForNextLevel(newLevel) {
		newLevel++
	}

	updateQuery := bson.M{
		"Experience": newExperience,
		"Level":      newLevel,
	}

	err := users.Patch(user.Key, updateQuery)
	if err != nil {
		return err
	}

	return nil
}

func (user *User) PercentToNextLevel() int {
	if user.Level == 1 {
		log.Println(100 * user.Experience / ExperienceForNextLevel(1))
		return 100 * user.Experience / ExperienceForNextLevel(1)
	}

	xp4level := experienceForLevel(user.Level)

	TotalXpForCurrentLevel :=
		ExperienceForNextLevel(user.Level) - xp4level

	UserXpForCurrentLevel := user.Experience - xp4level

	return 100 * UserXpForCurrentLevel / TotalXpForCurrentLevel
}

// experienceForLevel returns xp requiered to reach
// 'level'
func experienceForLevel(level int) int {
	if level <= 1 {
		return 0
	}
	// Divergent series: 1 + 2 + 3 + ... + level
	div := (level * (level + 1)) / 2
	return div * 16
}

// ExperienceForNextLevel returns the total experience
// points required to reach the level after 'level'
func ExperienceForNextLevel(level int) int {
	// Get the next level
	return experienceForLevel(level + 1)
}

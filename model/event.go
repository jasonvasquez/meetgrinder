package model

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/jasonvasquez/meetgrinder-api/types"

	"log"
)

const eventCollectionName = "events"

type dbOperation func(*mgo.Collection)

type Event struct {
	GenericObject
	Id           bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name         string        `json:"name"`
	Date         types.MGDate  `json:"date"`
	Heats        []Heat        `json:"heats"`
	Participants []Person      `json:"participants"`
}

func (evt *Event) Create() Event {

	now := bson.Now()
	evt.CreatedAt = now
	evt.UpdatedAt = now
	evt.Id = bson.NewObjectId()

	var newEvent Event
	dbPerform(func(c *mgo.Collection) {
		c.Insert(evt)
		log.Printf("New event: %v", evt)
		newEvent = *evt
	})

	return newEvent
}

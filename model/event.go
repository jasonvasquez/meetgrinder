package model

import (
	"github.com/jasonvasquez/meetgrinder-api/Godeps/_workspace/src/gopkg.in/mgo.v2"
	"github.com/jasonvasquez/meetgrinder-api/Godeps/_workspace/src/gopkg.in/mgo.v2/bson"

	"github.com/jasonvasquez/meetgrinder-api/types"

	"log"
)

const eventCollectionName = "events"

type Event struct {
	GenericObject
	Id           string       `json:"id" bson:"_id,omitempty"`
	Name         string       `json:"name" bson:"name"`
	Date         types.MGDate `json:"date" bson:"date"`
	Heats        []Heat       `json:"heats" bson:"heats"`
	Participants []Person     `json:"participants" bson:"participants"`
}

func FindAllEvents() []Event {

	var events []Event
	dbPerform(eventCollectionName, func(c *mgo.Collection) {
		c.Find(nil).Sort("date", "name").All(&events)
	})

	return events
}

func FindEventById(id string) interface{} {

	event := Event{}
	dbPerform(eventCollectionName, func(c *mgo.Collection) {
		log.Printf("Searching by id: %v", id)
		c.FindId(id).One(&event)
		log.Printf("Result: %v", event)
	})

	return event
}

func (evt *Event) Create() Event {

	now := bson.Now()
	evt.CreatedAt = now
	evt.UpdatedAt = now
	evt.Id = bson.NewObjectId().Hex()

	dbPerform(eventCollectionName, func(c *mgo.Collection) {
		c.Insert(evt)
	})

	return *evt
}

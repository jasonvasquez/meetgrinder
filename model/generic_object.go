package model

import (
	"time"
)

type GenericObject struct {
	Id        string    `json:"id"`
	Href      string    `json:"href"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FindById(i interface{}, id string) interface{} {

	// event := Event{Id: id, Name: "Event from finder"}
	// event := nil

	return nil
}

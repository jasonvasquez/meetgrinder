package types

import (
	"time"
)

type Event struct {
	GenericObject
	Id           string    `json:"id"`
	Name         string    `json:"name"`
	Date         time.Time `json:"date"`
	Heats        []Heat    `json:"heats"`
	Participants []Person  `jason:"participants"`
}

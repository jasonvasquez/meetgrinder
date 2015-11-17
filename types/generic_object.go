package types

import (
	"time"
)

type GenericObject struct {
	Id        string    `json:"id"`
	Href      string    `json:"href"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

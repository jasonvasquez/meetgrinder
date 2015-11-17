package model

type Person struct {
	GenericObject
	Name string `json:"name"`
}

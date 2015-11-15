package main

type Event struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Events []Event

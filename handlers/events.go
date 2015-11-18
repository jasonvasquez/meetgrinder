package handlers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/jasonvasquez/meetgrinder-api/model"
)

func EventIndex(w http.ResponseWriter, r *http.Request) {

	events := []model.Event{
		model.Event{Id: "1", Name: "Warren ES #1"},
		model.Event{Id: "2", Name: "Warren ES #2"},
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(events); err != nil {
		panic(err)
	}
}

func EventShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	eventId := vars["eventId"]

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var response = model.FindById(model.Event{}, eventId)

	if response == nil {
		w.WriteHeader(http.StatusNotFound)
		response = model.Error{http.StatusNotFound, "No Event found with that Id"}
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}

}

func EventCreate(w http.ResponseWriter, r *http.Request) {
	var event model.Event
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &event); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	// t := RepoCreateTodo(todo)
	t := event
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}

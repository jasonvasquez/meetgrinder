package main

import (
	"net/http"

	"github.com/jasonvasquez/meetgrinder-api/handlers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		handlers.Index,
	},
	Route{
		"EventIndex",
		"GET",
		"/events",
		handlers.EventIndex,
	},
	Route{
		"EventShow",
		"GET",
		"/events/{eventId}",
		handlers.EventShow,
	},
	Route{
		"EventCreate",
		"POST",
		"/events",
		handlers.EventCreate,
	},
}

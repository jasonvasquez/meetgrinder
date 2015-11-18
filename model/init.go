package model

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"os"
	"strings"
)

func init() {

	log.Println("Initializing Model Layer")
	var mongoURL string

	// Use a docker-style environment variable to locate the mongo instance
	// otherwise, fall back and assume localhost
	tcpURL := os.Getenv("DB_PORT")
	if tcpURL == "" {
		mongoURL = "mongodb://localhost"
	} else {
		mongoURL = strings.Replace(tcpURL, "tcp://", "mongodb://", -1)
	}

	log.Println("Connecting to mongo at", mongoURL)

	session, err := mgo.Dial(mongoURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("people")
	err = c.Insert(&Person{Name: "Ale"},
		&Person{Name: "Cla"})
	if err != nil {
		log.Fatal(err)
	}

	result := Person{}
	err = c.Find(bson.M{"name": "Ale"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Name:", result.Name)
}

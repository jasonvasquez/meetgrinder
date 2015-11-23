package model

import (
	"fmt"
	"github.com/jasonvasquez/meetgrinder-api/Godeps/_workspace/src/gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
	"log"
	// "os"
	// "strings"
)

const dbName = "meetgrinder"

var mongoSession mgo.Session

type logger struct{}

func (l logger) Output(calldepth int, s string) error {
	fmt.Printf("\t[%d]: %s\n", calldepth, s)
	return nil
}

func init() {
	log.Println("Initializing Model Layer")

	// mgo.SetDebug(true)
	mgo.SetLogger(logger{})

	// var mongoURL string

	// // Use a docker-style environment variable to locate the mongo instance
	// // otherwise, fall back and assume localhost
	// tcpURL := os.Getenv("DB_PORT")
	// if tcpURL == "" {
	// 	mongoURL = "mongodb://localhost"
	// } else {
	// 	mongoURL = strings.Replace(tcpURL, "tcp://", "mongodb://", -1)
	// }

	// log.Println("Connecting to mongo at", mongoURL)

	// mongoSession, err := mgo.Dial(mongoURL)
	// if err != nil {
	// 	panic(err)
	// }

	// // Optional. Switch the session to a monotonic behavior.
	// mongoSession.SetMode(mgo.Monotonic, true)

	// c := session.DB("test").C("people")
	// err = c.Insert(&Person{Name: "Ale"},
	// 	&Person{Name: "Cla"})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// result := Person{}
	// err = c.Find(bson.M{"name": "Ale"}).One(&result)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Name:", result.Name)
}

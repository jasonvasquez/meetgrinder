package model

import (
	"github.com/jasonvasquez/meetgrinder-api/Godeps/_workspace/src/gopkg.in/mgo.v2"

	"log"
	"os"
	"strings"
	"time"
)

type GenericObject struct {
	Id        string    `json:"id"`
	Href      string    `json:"href"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func dbPerform(collectionName string, op func(*mgo.Collection)) {
	// sess := mongoSession.Copy()
	// defer sess.Close()

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

	sess, err := mgo.Dial(mongoURL)
	if err != nil {
		panic(err)
	}
	defer sess.Close()

	// Optional. Switch the session to a monotonic behavior.
	sess.SetMode(mgo.Monotonic, true)

	collection := sess.DB(dbName).C(collectionName)
	op(collection)
}

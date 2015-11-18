package model

import (
	// "fmt"
	// "gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
	"log"
	// "os"
)

func init() {

	log.Println("Initializing Model Layer")

	// session, err := mgo.Dial("meetgrinder-db")
	// if err != nil {
	// 	panic(err)
	// }
	// defer session.Close()

	// // Optional. Switch the session to a monotonic behavior.
	// session.SetMode(mgo.Monotonic, true)

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

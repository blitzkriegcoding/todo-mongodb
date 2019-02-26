package config

import (
    "gopkg.in/mgo.v2"   
)

func InitDB(db_name string, collection string) (*mgo.Session, *mgo.Collection) {

    session, err := mgo.Dial("127.0.0.1")
    if err != nil {
        panic(err)
    }
    session.SetMode(mgo.Monotonic, true)

    // get a Collection of todo
    todosCollection := session.DB(db_name).C(collection)
    return session, todosCollection
}

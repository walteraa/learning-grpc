package models

import (
    "gopkg.in/mgo.v2/bson"
    "time"
)


type User struct {
    Id bson.ObjectId `bson:"_id json:id,omitempty"`
    FirstName   string
    LastName    string
    BirthDate   time.Time
}


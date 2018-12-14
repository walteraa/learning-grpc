package database

import (
    "log"
    "gopkg.in/mgo.v2"
)

type Session struct {
    session *mgo.Session
}


func NewSession(url string) (*Session, error){
    session, err := mgo.Dial(url)
    if err != nil {
        log.Fatal(err)
        return nil, err
    }

    return &Session{session}, err   
}   

func (s *Session) GetCollection(database string, collection string) *mgo.Collection{
    return s.session.DB(database).C(collection)
}

func (s *Session) Close(){
    if s.session != nil {
        s.session.Close()
    }
}

package database

import (
    "log"
    "sync"
    "gopkg.in/mgo.v2/bson"
    models "user_service/server/models"
)

// TODO: Use environment variables for database

const (
    COLLECTION_NAME = "users"
)

type UserDAO struct{
    Session *Session
}
var userDAO *UserDAO
var once sync.Once

func GetUserDAOInstance() (*UserDAO,error){

    var err error
    once.Do( func(){
        session, err := NewSession("localhost:27017")
        if err != nil {
            log.Fatal("Error when creating session")
            return
        }
        userDAO = &UserDAO{session}
    })
    return userDAO, err
}

func (ud *UserDAO) Insert(user models.User) error{
    err := userDAO.Session.GetCollection("myDB", COLLECTION_NAME).Insert(&user)
    return err
}

func (ud *UserDAO) FindById(id string) (models.User, error){
    var user models.User
    err := userDAO.Session.GetCollection( "myDB", COLLECTION_NAME).FindId(bson.ObjectIdHex(id)).One(&user)
    return user, err
}

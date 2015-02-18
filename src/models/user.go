package models

import (
	"code.google.com/p/go.crypto/bcrypt"
	"github.com/golang/glog"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type User struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Email     string        `bson:"email"`
	Username  string        `bson:"username"`
	Password  []byte        `bson:"psswrd"`
	Timestamp time.Time     `bson:"timestamp"`
}

func (user *User) HashPassword(password string) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		glog.Fatalf("Couldn't hash password: %v", err)
		panic(err)
	}
	user.Password = hash
}

func GetUserByEmail(database *mgo.Database, email string) (user *User) {
	err := database.C("users").Find(bson.M{"email": email}).One(&user)

	if err != nil {
		glog.Warningf("Can't get user by email: %v", err)
	}
	return
}

func GetUserById(database *mgo.Database, id bson.ObjectId) (user *User) {
	err := database.C("users").Find(bson.M{"_id": id}).One(&user)

	if err != nil {
		glog.Warningf("Can't get user by email: %v", err)
	}
	return
}

func InsertUser(database *mgo.Database, user *User) error {
	user.ID = bson.NewObjectId()
	return database.C("users").Insert(user)
}

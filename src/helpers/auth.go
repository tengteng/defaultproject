package helpers

import (
	"models"

	"code.google.com/p/go.crypto/bcrypt"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func Login(c *mgo.Database, email string, password string) (user *models.User, err error) {
	err = c.C("users").Find(bson.M{"email": email}).One(&user)
	if err != nil {
		return
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	if err != nil {
		user = nil
	}
	return
}

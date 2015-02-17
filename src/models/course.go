package models

import (
	"github.com/golang/glog"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Course struct {
	ID           bson.ObjectId `bson:"_id,omitempty"`
	Creator      bson.ObjectId `bson:"_creator_id,omitempty"`
	CreatorEmail string        `bson:"creator_email"`
	Timestamp    time.Time     `bson:"timestamp"`
	Title        string        `bson:"title"`
}

func GetCourseByID(database *mgo.Database, id bson.ObjectId) (course *Course) {
	err := database.C("courses").Find(bson.M{"_id": id}).One(&course)

	if err != nil {
		glog.Warningf("Can't get user by email: %v", err)
	}
	return
}

func InsertCourse(database *mgo.Database, course *Course) error {
	course.ID = bson.NewObjectId()
	return database.C("courses").Insert(course)
}

package models

import (
	"github.com/golang/glog"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Course struct {
	ID              bson.ObjectId `bson:"_id,omitempty"`
	Creator         bson.ObjectId `bson:"_creator_id,omitempty"`
	CreatorEmail    string        `bson:"creator_email"`
	Timestamp       time.Time     `bson:"timestamp"`
	CourseTitle     string        `bson:"course_title"`
	Description     string        `bson:"description"`
	Tuition         float64       `bson:"tuition"`
	MaxParticipants int64         `bson:"max_participants"`
	ExpireDate      time.Time     `bson:"expire_date"`
	TeachingMethod  int           `bson:"teaching_method"`
}

func GetCourseByID(database *mgo.Database, id bson.ObjectId) (course *Course) {
	err := database.C("courses").Find(bson.M{"_id": id}).One(&course)

	if err != nil {
		glog.Warningf("Can't get user by email: %v", err)
	}
	return
}

func GetNCourseByCreatorId(database *mgo.Database, id bson.ObjectId,
	N int) (courses []*Course) {
	if N > 0 {
		err := database.C("courses").Find(
			bson.M{"_id": id}).Limit(N).All(&courses)

		if err != nil {
			glog.Warningf("Can't get user by email: %v", err)
		}
	} else {
		// N <=0 , get all courses.
		err := database.C("courses").Find(
			bson.M{"_creator_id": id}).All(&courses)

		if err != nil {
			glog.Warningf("Can't get user by email: %v", err)
		}
	}
	return
}

func InsertCourse(database *mgo.Database, course *Course) error {
	course.ID = bson.NewObjectId()
	return database.C("courses").Insert(course)
}

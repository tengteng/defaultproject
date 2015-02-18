package web

import (
	"net/http"
	"strconv"
	"time"
	// "html/template"

	"helpers"
	"models"

	"github.com/golang/glog"
	"github.com/zenazn/goji/web"
	"gopkg.in/mgo.v2/bson"
)

func (controller *Controller) LaunchCourse(c web.C, r *http.Request) (string, int) {
	session := controller.GetSession(c)
	if session.Values["User"] == nil {
		// SignIn first.
		return "/signin", http.StatusSeeOther
	}

	t := controller.GetTemplate(c)
	return helpers.Parse(t, "createcourse", nil), http.StatusOK
}

func (controller *Controller) LaunchCoursePost(c web.C, r *http.Request) (string, int) {
	course_title := r.FormValue("courseTitle")
	course_description := r.FormValue("courseDescription")
	tuition, _ := strconv.ParseFloat(r.FormValue("tuition"), 64)
	max_participants, _ := strconv.Atoi(r.FormValue("max_participants"))
	teaching_method := r.FormValue("teaching_method")
	const shortForm = "2006-01-02"
	expire, _ := time.Parse(shortForm, r.FormValue("expire_date"))

	session := controller.GetSession(c)
	database := controller.GetDatabase(c)
	course := &models.Course{
		Creator:         session.Values["User"].(bson.ObjectId),
		CreatorEmail:    session.Values["UserEmail"].(string),
		Timestamp:       time.Now(),
		CourseTitle:     course_title,
		Description:     course_description,
		Tuition:         tuition,
		MaxParticipants: int64(max_participants),
		ExpireDate:      expire,
		TeachingMethod:  teaching_method,
	}

	if err := models.InsertCourse(database, course); err != nil {
		session.AddFlash("Error while creating course.")
		glog.Errorf("Error whilst creating course: %v", err)
		return controller.LaunchCourse(c, r)
	}

	return "/", http.StatusSeeOther
}

// func validateCourse() bool {
//
// }

package web

import (
	"fmt"
	"net/http"
	// "reflect"
	"strconv"
	"time"

	"helpers"
	"models"

	"github.com/golang/glog"
	"github.com/zenazn/goji/web"
	"gopkg.in/mgo.v2/bson"
	// "html/template"
)

func (controller *Controller) LaunchCourse(c web.C, r *http.Request) (string, int) {
	session := controller.GetSession(c)
	if session.Values["User"] == nil {
		// SignIn first.
		return "/signin", http.StatusSeeOther
	}

	t := controller.GetTemplate(c)
	// widgets := helpers.Parse(t, "createcourse", nil)
	return helpers.Parse(t, "createcourse", nil), http.StatusOK

	// return "/", http.StatusSeeOther
}

func (controller *Controller) LaunchCoursePost(c web.C, r *http.Request) (string, int) {
	// fmt.Println("TUITION: ", r.FormValue("tuition"))
	// fmt.Println("COURSEDESCRIPTION: ", r.FormValue("courseDescription"))
	// fmt.Println("TERMS: ", r.FormValue("terms"))
	// fmt.Println("COURSETITLE: ", r.FormValue("courseTitle"))
	// fmt.Println("DESCRIPTION: ", r.FormValue("courseDescription"))
	// fmt.Println("TUITION: ", r.FormValue("tuition"))
	// fmt.Println("MAX_PARTICIPANTS: ", r.FormValue("max_participants"))
	// fmt.Println("EXPIRE_DATE: ", r.FormValue("expire_date"))
	// fmt.Println("TEACHINGMETHOD: ", r.FormValue("teaching_method"))
	// fmt.Println("TEACHINGMETHOD: ", reflect.TypeOf(r.FormValue("teaching_method")))

	/*
	   TUITION:  102
	   COURSEDESCRIPTION:  golang desc blablabla
	   TERMS:  on
	   COURSETITLE:  golang abc
	   DESCRIPTION:  golang desc blablabla
	   TUITION:  102
	   MAX_PARTICIPANTS:  20
	   EXPIRE_DATE:  2015-02-25
	   TEACHINGMETHOD:  Online
	*/
	course_title := r.FormValue("courseTitle")
	course_description := r.FormValue("courseDescription")
	tuition, _ := strconv.ParseFloat(r.FormValue("tuition"), 64)
	max_participants, _ := strconv.Atoi(r.FormValue("max_participants"))
	teaching_method := r.FormValue("teaching_method")
	const shortForm = "2006-01-02"
	expire, _ := time.Parse(shortForm, r.FormValue("expire_date"))

	session := controller.GetSession(c)
	database := controller.GetDatabase(c)

	fmt.Println(session.Values["User"].(bson.ObjectId))
	fmt.Println(session.Values["Email"].(string))
	fmt.Println(time.Now())
	fmt.Println(course_title)
	fmt.Println(course_description)
	fmt.Println(tuition)
	fmt.Println(int64(max_participants))
	fmt.Println(expire)
	fmt.Println(teaching_method)

	course := &models.Course{
		Creator:         session.Values["User"].(bson.ObjectId),
		CreatorEmail:    session.Values["Email"].(string),
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

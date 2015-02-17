package web

import (
	"fmt"
	"net/http"

	"helpers"
	// "models"

	"github.com/zenazn/goji/web"
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
	fmt.Println("TUITION: ", r.FormValue("tuition"))
	fmt.Println("COURSEDESCRIPTION: ", r.FormValue("courseDescription"))
	fmt.Println("TERMS: ", r.FormValue("terms"))
	fmt.Println("COURSETITLE: ", r.FormValue("courseTitle"))
	fmt.Println("DESCRIPTION: ", r.FormValue("courseDescription"))
	fmt.Println("TUITION: ", r.FormValue("tuition"))
	fmt.Println("MAX_PARTICIPANTS: ", r.FormValue("max_participants"))
	fmt.Println("EXPIRE_DATE: ", r.FormValue("expire_date"))
	fmt.Println("TEACHINGMETHOD: ", r.FormValue("teaching_method"))

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

	/*
	   const shortForm = "2006-01-02"
	   t, _ = time.Parse(shortForm, "2013-02-03")
	*/

	/*
	   session := controller.GetSession(c)
	   database := controller.GetDatabase(c)
	   course := &models.Course{
	       Creator: ,
	       ID: ,
	       Creator: ,
	       CreatorEmail: ,
	       Timestamp: ,
	       CourseTitle: ,
	       Description: ,
	       Tuition: ,
	       MaxParticipants: ,
	       ExpireDate: ,
	       TeachingMethod:
	   }

	   if err := models.InsertCourse(database, course); err != nil {
	       session.AddFlash("Error while creating course.")
	       glog.Errorf("Error whilst creating course: %v", err)
	       return controller.LaunchCourse(c, r)
	   }
	*/

	return "/", http.StatusSeeOther
}

// func validateCourse() bool {
//
// }

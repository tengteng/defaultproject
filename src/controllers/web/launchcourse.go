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

	/*
			session := controller.GetSession(c)
			database := controller.GetDatabase(c)
			course := &models.Course{
		            Creator: ,

		        }
	*/

	return "/", http.StatusSeeOther
}

// func validateCourse() bool {
//
// }

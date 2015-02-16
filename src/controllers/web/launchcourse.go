package web

import (
	"net/http"

	"helpers"

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

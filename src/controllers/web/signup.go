package web

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"models"

	"github.com/golang/glog"
	"github.com/zenazn/goji/web"
)

// Sign up route
func (controller *Controller) SignUp(c web.C, r *http.Request) (string, int) {

	fmt.Println("XXXXX: ", c)

	t := controller.GetTemplate(c)
	session := controller.GetSession(c)

	// With that kind of flags template can "figure out" what route is being rendered
	c.Env["IsSignUp"] = true

	c.Env["Flash"] = session.Flashes("auth")

	var widgets = controller.Parse(t, "auth/signup", c.Env)

	c.Env["Title"] = "Default Project - Sign Up"
	c.Env["Content"] = template.HTML(widgets)

	return controller.Parse(t, "main", c.Env), http.StatusOK
}

// Sign Up form submit route. Registers new user or shows Sign Up route with appropriate messages set in session
func (controller *Controller) SignUpPost(c web.C, r *http.Request) (string, int) {
	email, password := r.FormValue("email"), r.FormValue("password")

	session := controller.GetSession(c)
	database := controller.GetDatabase(c)

	user := models.GetUserByEmail(database, email)

	if user != nil {
		session.AddFlash("User exists", "auth")
		return controller.SignUp(c, r)
	}

	user = &models.User{
		Username:  email,
		Email:     email,
		Timestamp: time.Now(),
	}
	user.HashPassword(password)

	if err := models.InsertUser(database, user); err != nil {
		session.AddFlash("Error while registering user.")
		glog.Errorf("Error while registering user: %v", err)
		return controller.SignUp(c, r)
	}

	session.Values["User"] = user.ID

	return "/", http.StatusSeeOther
}

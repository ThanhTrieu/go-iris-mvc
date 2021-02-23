package controllers

import (
	"gomvc/services"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
)

type DashboardController struct {
	Ctx     iris.Context
	Service services.UsersService
	Session *sessions.Session
}

const IDKey = "BACKEND_LOGGING_1990"

func (c *DashboardController) getCurrentUserID() int64 {
	userID := c.Session.GetInt64Default(IDKey, 0)
	return userID
}

func (c *DashboardController) getCurrentUsername() string {
	username := c.Session.GetStringDefault("usernameSession", "")
	return username
}

func (c *DashboardController) isLoggedIn() bool {
	u := c.getCurrentUsername()
	id := c.getCurrentUserID()
	if u == "" || id <= 0 {
		return false
	}
	return true
}

var dashboardStaticView = mvc.View {
	Name: "dashboard/index.html",
	Data: iris.Map{
		"Title": "Dashboard page",
		"layout": true,
	},
}
var LoginStaticPath =  mvc.Response { 
	Path: "/user/login",
}

func (c *DashboardController) GetDashboard() mvc.Result {
	if c.isLoggedIn() {
		return dashboardStaticView
	}
	return LoginStaticPath
}

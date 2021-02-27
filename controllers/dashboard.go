package controllers

import (
	"gomvc/helpers"
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
	if helpers.IsLoggedIn(c.Ctx) {
		return dashboardStaticView
	}
	return LoginStaticPath
}

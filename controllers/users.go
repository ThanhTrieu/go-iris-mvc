package controllers

import (
	"gomvc/services"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
)

type UsersController struct {
	Ctx     iris.Context
	Service services.UsersService
	Session *sessions.Session
}

const userIDKey = "BACKEND_LOGGING_1990"

func (c *UsersController) getCurrentUserID() int64 {
	userID := c.Session.GetInt64Default(userIDKey, 0)
	return userID
}

func (c *UsersController) isLoggedIn() bool {
	return c.getCurrentUserID() > 0
}

func (c *UsersController) PostLogout() mvc.Result {
	c.Session.Destroy()
	return mvc.Response{
		Path: "/user/login",
	}
}

var loginUserStaticView = mvc.View{
	Name: "users/login.html",
	Data: iris.Map{"Title": "User login page"},
}

//http://localhost:8080/user/login
func (c *UsersController) GetLogin() mvc.Result {
	if c.isLoggedIn() {
		return mvc.Response{
			Path: "/admin/dashboard",
		}
	}

	return loginUserStaticView
}

func (c *UsersController) PostLogin() mvc.Result {
	var (
		username = c.Ctx.FormValue("username")
		password = c.Ctx.FormValue("password")
	)

	u, check := c.Service.CheckLoginUser(username, password)

	if check == false {
		return mvc.Response{
			Path: "/user/login",
		}
	}

	c.Session.Set(userIDKey, u.ID)
	c.Session.Set(userIDKey, u.Username)
	c.Session.Set(userIDKey, u.Email)
	c.Session.Set(userIDKey, u.AuthenKey)
	c.Session.Set(userIDKey, u.Phone)
	c.Session.Set(userIDKey, u.GroupId)

	return mvc.Response{
		Path: "/admin/dashboard",
	}
}

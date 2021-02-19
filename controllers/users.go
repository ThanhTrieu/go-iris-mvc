package controllers

import (
	"fmt"
	"gomvc/services"
	"gomvc/requests"
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

//http://localhost:8080/user/login
func (c *UsersController) GetLogin() mvc.Result {
	if c.isLoggedIn() {
		return mvc.Response{
			Path: "/admin/dashboard",
		}
	}
	errs := c.Ctx.URLParam("state")

	return mvc.View {
		Name: "users/login.html",
		Data: iris.Map{
			"Title": "User login page",
			"message": errs,
		},
	}
}

func (c *UsersController) PostLogin() mvc.Result {
	var (
		username = c.Ctx.FormValue("username")
		password = c.Ctx.FormValue("password")
	)

	u, check := c.Service.CheckLoginUser(username, password)

	if check == false {
		return mvc.Response{
			Path: "/user/login?state=fail",
		}
	}

	c.Session.Set(userIDKey, u.ID)
	c.Session.Set(userIDKey, u.Username)
	c.Session.Set(userIDKey, u.Email)
	c.Session.Set(userIDKey, u.AuthenKey)
	c.Session.Set(userIDKey, u.Phone)

	return mvc.Response{
		Path: "/admin/dashboard",
	}
}

func (c *UsersController) GetRegister() mvc.Result {
	if c.isLoggedIn() {
		return mvc.Response{
			Path: "/admin/dashboard",
		}
	}

	return mvc.View {
		Name: "users/register.html",
		Data: iris.Map{
			"Title": "User register page",
		},
	}
}

func (c *UsersController) PostRegister() mvc.Result {
	var (
		username = c.Ctx.FormValue("username")
		password = c.Ctx.FormValue("password")
		email = c.Ctx.FormValue("email")
		phone = c.Ctx.FormValue("phone")
	)
	msg := &requests.Message{
		Username: username,
		Password: password,
		Phone:  phone,
		Email:   email,
	}
	if msg.Validate() == false {
		fmt.Println(msg.Errors)
		return mvc.Response {
			Path: "/user/register",
		}
	}
	return mvc.Response{
		Path: "/user/login",
	}
}

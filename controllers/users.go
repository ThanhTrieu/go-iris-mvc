package controllers

import (
	"gomvc/requests"
	"gomvc/services"
	"html"
	"os"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
)

const userIDKey = "BACKEND_LOGGING_1990"

type UsersController struct {
	Ctx     iris.Context
	Service services.UsersService
	Session *sessions.Session
}

func (c *UsersController) getCurrentUserID() int64 {
	userID := c.Session.GetInt64Default(userIDKey, 0)
	return userID
}

func (c *UsersController) getCurrentUsername() string {
	username := c.Session.GetStringDefault("usernameSession", "")
	return username
}

func (c *UsersController) isLoggedIn() bool {
	u := c.getCurrentUsername()
	if u == "" {
		return false
	}
	return true
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

	dataSiteKey := os.Getenv("DATA_SITE_KEY")
	errs := c.Ctx.URLParam("state")
	return mvc.View {
		Name: "users/login.html",
		Data: iris.Map{
			"Title": "User login page",
			"message": errs,
			"siteKey": dataSiteKey,
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
	
	c.Session.Set(userIDKey, int64(u.ID))
	c.Session.Set("usernameSession", u.Username)
	c.Session.Set("emailSession", u.Email)
	c.Session.Set("authenSession", u.AuthenKey)
	c.Session.Set("phoneSession", u.Phone)

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
	dataSiteKey := os.Getenv("DATA_SITE_KEY")
	msg := c.Session.Get("ErrorsRegister")
	errUser := c.Session.Get("ErrorsUsername")
	errEmail := c.Session.Get("ErrorsEmail")

	state := c.Ctx.URLParam("state")
	if state != "fail" || state != "error_user" || state != "error_email" {
		c.Session.Destroy()
	} 

	return mvc.View {
		Name: "users/register.html",
		Data: iris.Map{
			"Title": "User register page",
			"msg": msg,
			"errUser": errUser,
			"errEmail": errEmail,
			"siteKey": dataSiteKey,
		},
	}
}

func (c *UsersController) PostRegister() mvc.Result {
	var (
		username = html.EscapeString(c.Ctx.FormValue("username"))
		password = html.EscapeString(c.Ctx.FormValue("password"))
		email = html.EscapeString(c.Ctx.FormValue("email"))
		phone = html.EscapeString(c.Ctx.FormValue("phone"))
	)
	msg := &requests.Message{
		Username: username,
		Password: password,
		Phone:  phone,
		Email:   email,
	}
	if msg.Validate() == false {
		c.Session.Set("ErrorsRegister", msg.Errors)
		return mvc.Response {
			Path: "/user/register?state=fail",
		}
	} 
	
	errUser := c.Service.CheckUsernameExists(username)
	if errUser {
		c.Session.Set("ErrorsUsername", "Username exists")
		return mvc.Response {
			Path: "/user/register?state=error_user",
		}
	}
	errEmail := c.Service.CheckEmailExists(email)
	if errEmail {
		c.Session.Set("ErrorsEmail", "Email exists")
		return mvc.Response {
			Path: "/user/register?state=error_email",
		}
	}

	c.Session.Destroy()
	insert := c.Service.CreateUser(username, password, email, phone)
	if insert == false {
		return mvc.Response {
			Path: "/user/register?state=error",
		}
	} 
	return mvc.Response {
		Path: "/user/login?mess=register_success",
	}
}

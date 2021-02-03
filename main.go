package main

import (
	"gomvc/controllers"
	"gomvc/database"
	"gomvc/models"
	"gomvc/repos"
	"gomvc/services"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")

	//masterpage
	tmpl := iris.HTML("./templates", ".html").Layout("masterpage.html").Reload(true)
	app.RegisterView(tmpl)

	//static files
	app.HandleDir("/static", iris.Dir("./public/static"))

	//routes
	app.Get("/", func(ctx iris.Context) {
		ctx.Writef("Hello world")
	})

	// **** (MySQL)
	db, err := database.ConnectSQL("127.0.0.1", "3306", "root", "", "go_lang")
	if err != nil {
		app.Logger().Fatalf("error while loading the tables: %v", err)
		return
	}
	//for migrate
	db.AutoMigrate(&models.Users{})

	// session
	sessManager := sessions.New(sessions.Config{
		Cookie:  "IueBm5pJGVe5dzsQ",
		Expires: 24 * time.Hour,
	})

	//for user
	userRepo := repos.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	users := mvc.New(app.Party("/user"))
	users.Register(
		userService,
		sessManager.Start,
	)
	users.Handle(new(controllers.UsersController))

	// for dashboard
	dashboardRepo := repos.NewUserRepository(db)
	dashboardService := services.NewUserService(dashboardRepo)

	dashboards := mvc.New(app.Party("/admin"))
	dashboards.Register(
		dashboardService,
		sessManager.Start,
	)
	dashboards.Handle(new(controllers.DashboardController))

	//error
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("Message", ctx.Values().GetStringDefault("message", "The page you're looking for doesn't exist"))
		ctx.View("error.html")
	})

	//start
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}

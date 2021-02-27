package main

import (
	"gomvc/controllers"
	"gomvc/database"
	"gomvc/models"
	"gomvc/repos"
	"gomvc/services"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
)

func setSessionViewData(ctx iris.Context) {
	session := sessions.Get(ctx)
	username := session.GetStringDefault("usernameSession", "")
	roleUser := session.GetInt64Default("roleSession", 0)
	ctx.ViewData("username", username)
	ctx.ViewData("roleUser", roleUser)
	ctx.Next()
}

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	errEnv := godotenv.Load()
	if errEnv != nil {
    app.Logger().Fatalf("Error loading .env file: %v", errEnv)
  }

	//masterpage
	tmpl := iris.HTML("./templates", ".html").Layout("masterpage.html").Reload(true)
	app.RegisterView(tmpl)

	//static files
	app.HandleDir("/static", iris.Dir("./public/static"))

	// **** (MySQL)
	db, err := database.ConnectSQL(
		os.Getenv("DB_HOST_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))

	if err != nil {
		app.Logger().Fatalf("error while loading the tables: %v", err)
		return
	}

	// close db
	sqlDB, errDb := db.DB()
	if errDb != nil {
		app.Logger().Fatalf("can not close database: %v", errDb)
		return
	}
	defer sqlDB.Close()

	//for migrate
	db.AutoMigrate(&models.Users{}, &models.Groups{}, &models.LeaderFolders{}, &models.MemberFolder{})

	// session
	sessManager := sessions.New(sessions.Config{
		Cookie:  "IueBm5pJGVe5dzsQ",
		Expires: 24 * time.Hour,
	})
	// using session in view
	app.Use(sessManager.Handler())
	app.Use(setSessionViewData)

	// csrf token
	/*
	CSRF := csrf.Protect(
		// Note that the authentication key
		// provided should be 32 bytes
		// long and persist across application restarts.
		[]byte("9AB0F421E53A477C084477AEA06096F5"),
		// WARNING: Set it to true on production with HTTPS.
		csrf.Secure(false),
	)
	*/
	//for user
	userRepo := repos.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	users := mvc.New(app.Party("/user"))
	users.Register(
		userService,
		sessManager.Start,
	)
	// users.Use(CSRF)
	users.Handle(new(controllers.UsersController))

	// for dashboard
	dashboardRepo := repos.NewGroupRepository(db)
	dashboardService := services.NewGroupService(dashboardRepo)

	dashboards := mvc.New(app.Party("/admin"))
	dashboards.Register(
		dashboardService,
		sessManager.Start,
	)
	dashboards.Handle(new(controllers.DashboardController))

	// for create folder leader and member
	leaderFolderRepo := repos.NewLeaderFolderRepository(db)
	leaderFolderService := services.NewLeaderFoldersService(leaderFolderRepo)

	memberFolderRepo := repos.NewMemberFolderRepository(db)
	memberFolderService := services.NewMemberFoldersService(memberFolderRepo)

	groupRepo := repos.NewGroupRepository(db)
	groupService := services.NewGroupService(groupRepo)

	fileListFolderRepo := repos.NewFileListFolderRepository(db)
	fileListFolderService := services.NewFileListFoldersService(fileListFolderRepo)

	createFolder := mvc.New(app.Party("/admin"))
	createFolder.Register(
		groupService,
		leaderFolderService,
		memberFolderService,
		fileListFolderService,
		sessManager.Start,
	)
	createFolder.Handle(new(controllers.FolderController))

	//error
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("Message", ctx.Values().GetStringDefault("message", "The page you're looking for doesn't exist"))
		ctx.View("error.html")
	})

	//start
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}

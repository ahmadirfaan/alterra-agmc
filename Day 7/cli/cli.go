package cli

import (
	"alterra-agmc-day7/app"
	database_config "alterra-agmc-day7/config/database"
	controller "alterra-agmc-day7/controllers"
	"alterra-agmc-day7/middleware"
	"alterra-agmc-day7/repositories"
	"alterra-agmc-day7/services"
	"github.com/labstack/echo/v4"
)

type Cli struct {
	Args []string
}

func NewCli(args []string) *Cli {
	return &Cli{
		Args: args,
	}
}

func (cli *Cli) Run(application *app.Application) {
	// set up connection
	db := database_config.InitDb()

	//Repository
	bookRepo := repositories.NewBookRepository(db)
	userRepo := repositories.NewUserRepository(db)

	// Service
	bookService := services.NewBookService(bookRepo)
	userService := services.NewUserService(userRepo)

	// Controller
	bookController := controller.NewBookController(bookService)
	userController := controller.NewUserController(userService)

	e := echo.New()

	middleware.LogMiddleware(e)
	e.Validator = middleware.NewCustomValidator()

	//Controller for users
	e.POST("/v1/users", userController.CreateUser)
	e.POST("/v1/login", userController.LoginUser)

	//Controller for book
	e.GET("/v1/books/:id", bookController.GetBookById)
	e.GET("/v1/books", bookController.GetAllBooks)

	//restricted path
	groupJWT := e.Group("/restricted")
	middleware.SetJwtMiddlewares(groupJWT)
	groupJWT.POST("/v1/books", bookController.CreateBook)
	groupJWT.PUT("/v1/books/:id", bookController.UpdateBook)
	groupJWT.DELETE("/v1/books/:id", bookController.DeleteBook)

	groupJWT.PUT("/v1/users/:id", userController.UpdateUser)
	groupJWT.GET("/v1/users/:id", userController.GetUserById)
	groupJWT.GET("/v1/users", userController.GetAllUsers)
	groupJWT.DELETE("/v1/users/:id", userController.DeleteUser)

	e.Logger.Fatal(e.Start(":" + application.Config.AppPort))
}

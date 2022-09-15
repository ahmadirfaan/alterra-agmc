package cli

import (
	"alterra-agmc-dynamic-crud/app"
	database_config "alterra-agmc-dynamic-crud/config/database"
	controller "alterra-agmc-dynamic-crud/controllers"
	"alterra-agmc-dynamic-crud/repositories"
	"alterra-agmc-dynamic-crud/services"
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

	//Controller for users
	e.POST("/v1/users", userController.CreateUser)
	e.PUT("/v1/users/:id", userController.UpdateUser)
	e.GET("/v1/users/:id", userController.GetUserById)
	e.GET("/v1/users", userController.GetAllUsers)
	e.DELETE("/v1/users/:id", userController.DeleteUser)

	//Controller for book
	e.POST("/v1/books", bookController.CreateBook)
	e.GET("/v1/books/:id", bookController.GetBookById)
	e.GET("/v1/books", bookController.GetAllBooks)
	e.PUT("/v1/books/:id", bookController.UpdateBook)
	e.DELETE("/v1/books/:id", bookController.DeleteBook)

	e.Logger.Fatal(e.Start(":" + application.Config.AppPort))
}

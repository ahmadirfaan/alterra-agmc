package main

import (
	"github.com/labstack/echo/v4"
	"static-crud/controller"
)

func main() {
	e := echo.New()

	v1 := e.Group("/v1")
	books := v1.Group("/books")

	books.GET("", controller.GetAllBooks)
	books.GET("/:id", controller.GetBookById)
	books.POST("", controller.CreateNewBook)
	books.PUT("/:id", controller.UpdateBookById)
	books.DELETE("/:id", controller.DeleteBookById)

	e.Logger.Fatal(e.Start(":8080"))
}

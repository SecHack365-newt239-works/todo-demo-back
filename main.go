package main

import (
	"net/http"
	"todo-demo-back/routes/todo"

	"github.com/labstack/echo/v4"
)

func Hello() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	}
}

func main() {
	e := echo.New()
	e.GET("/", Hello())
	e.GET("/todo", todo.GetTodoByID())
	e.Logger.Fatal(e.Start(":1323"))
}

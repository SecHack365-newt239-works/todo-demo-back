package main

import (
	"net/http"
	"todo-demo-back/routes/todo"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Hello() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.GET("/", Hello())
	e.GET("/todo", todo.GetTodoByID())
	e.POST("/todo", todo.CreateTodo())
	e.PUT("/todo/:id", todo.UpdateTodoDone())
	e.Logger.Fatal(e.Start(":1323"))
}

package main

import (
	"log"
	"net/http"
	"os"
	"todo-demo-back/routes/todo"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Hello() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	}
}

func main() {
	if os.Getenv("ENV") == "local" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8000", os.Getenv("FRONTEND_URL")},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.GET("/", Hello())
	e.GET("/todo", todo.GetTodo())
	e.POST("/todo", todo.CreateTodo())
	e.PUT("/todo/:id", todo.DeleteTodo())
	e.Logger.Fatal(e.Start(":1323"))
}
